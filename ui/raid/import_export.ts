import { Exporter } from '/tbc/core/components/exporters.js';
import { Importer } from '/tbc/core/components/importers.js';
import { RaidSimSettings } from '/tbc/core/proto/ui.js';
import { Party, Player, Raid, TargetedActionMetrics } from '../core/proto/api.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';

import { RaidSimUI } from './raid_sim_ui.js';
import { Encounter, EquipmentSpec, Gem, ItemSpec, MobType, Spec, Target } from '../core/proto/common.js';
import { nameToClass } from '../core/proto_utils/names.js';
import { Faction, raceToFaction, specToClass, specToEligibleRaces, specTypeFunctions, withSpecProto } from '../core/proto_utils/utils.js';
import { BalanceDruid, BalanceDruid_Rotation_PrimarySpell, FeralDruid } from '../core/proto/druid.js';
import { ElementalShaman, EnhancementShaman } from '../core/proto/shaman.js';
import { Hunter } from '../core/proto/hunter.js';
import { Mage } from '../core/proto/mage.js';
import { Rogue } from '../core/proto/rogue.js';
import { RetributionPaladin } from '../core/proto/paladin.js';
import { ShadowPriest, SmitePriest } from '../core/proto/priest.js';
import { Warlock } from '../core/proto/warlock.js';
import { ProtectionWarrior, Warrior } from '../core/proto/warrior.js';
import { gemMatchesSocket } from '../core/proto_utils/gems.js';
import { playerPresets } from './presets.js';

declare var $: any;
declare var tippy: any;

export function newRaidImporters(simUI: RaidSimUI): HTMLElement {
	const importSettings = document.createElement('div');
	importSettings.classList.add('import-settings', 'sim-dropdown-menu');
	importSettings.innerHTML = `
		<span id="importMenuLink" class="dropdown-toggle fas fa-file-import" role="button" data-toggle="dropdown" aria-haspopup="true" arai-expanded="false"></span>
		<div class="dropdown-menu dropdown-menu-right" aria-labelledby="importMenuLink">
		</div>
	`;
	const linkElem = importSettings.getElementsByClassName('dropdown-toggle')[0] as HTMLElement;
	tippy(linkElem, {
		'content': 'Import',
		'allowHTML': true,
	});

	const menuElem = importSettings.getElementsByClassName('dropdown-menu')[0] as HTMLElement;
	const addMenuItem = (label: string, onClick: () => void) => {
		const itemElem = document.createElement('span');
		itemElem.classList.add('dropdown-item');
		itemElem.textContent = label;
		itemElem.addEventListener('click', onClick);
		menuElem.appendChild(itemElem);
	};

	addMenuItem('Json', () => new RaidJsonImporter(menuElem, simUI));
	addMenuItem('WCL', () => new RaidWCLImporter(menuElem, simUI));

	return importSettings;
}

export function newRaidExporters(simUI: RaidSimUI): HTMLElement {
	const exportSettings = document.createElement('div');
	exportSettings.classList.add('export-settings', 'sim-dropdown-menu');
	exportSettings.innerHTML = `
		<span id="exportMenuLink" class="dropdown-toggle fas fa-file-export" role="button" data-toggle="dropdown" aria-haspopup="true" arai-expanded="false"></span>
		<div class="dropdown-menu dropdown-menu-right" aria-labelledby="exportMenuLink">
		</div>
	`;
	const linkElem = exportSettings.getElementsByClassName('dropdown-toggle')[0] as HTMLElement;
	tippy(linkElem, {
		'content': 'Export',
		'allowHTML': true,
	});

	const menuElem = exportSettings.getElementsByClassName('dropdown-menu')[0] as HTMLElement;
	const addMenuItem = (label: string, onClick: () => void) => {
		const itemElem = document.createElement('span');
		itemElem.classList.add('dropdown-item');
		itemElem.textContent = label;
		itemElem.addEventListener('click', onClick);
		menuElem.appendChild(itemElem);
	};

	addMenuItem('Json', () => new RaidJsonExporter(menuElem, simUI));

	return exportSettings;
}

class RaidJsonImporter extends Importer {
	private readonly simUI: RaidSimUI;
	constructor(parent: HTMLElement, simUI: RaidSimUI) {
		super(parent, 'JSON Import');
		this.simUI = simUI;

		this.descriptionElem.innerHTML = `
			<p>
				Import settings from a JSON text file, which can be created using the JSON Export feature of this site.
			</p>
			<p>
				To import, paste the JSON text below and click, 'Import'.
			</p>
		`;
	}

	onImport(data: string) {
		const settings = RaidSimSettings.fromJsonString(data);
		this.simUI.fromProto(TypedEvent.nextEventID(), settings);
		this.close();
	}
}

class RaidJsonExporter extends Exporter {
	private readonly simUI: RaidSimUI;

	constructor(parent: HTMLElement, simUI: RaidSimUI) {
		super(parent, 'JSON Export', true);
		this.simUI = simUI;
		this.init();
	}

	getData(): string {
		return JSON.stringify(RaidSimSettings.toJson(this.simUI.toProto()), null, 2);
	}
}

class RaidWCLImporter extends Importer {
	private readonly simUI: RaidSimUI;
	constructor(parent: HTMLElement, simUI: RaidSimUI) {
		super(parent, 'WCL Import');
		this.simUI = simUI;
		this.descriptionElem.innerHTML = `
			<p>
				Import entire raid from a WCL report
			</p>
			<p>
				To import, paste the WCL report and fight link (https://classic.warcraftlogs.com/reports/REPORTID#fight=FIGHTID)
			</p>
		`;
	}

	onImport(importLink: string) {
		// TODO: validate link so we dont get a crash
		if (!importLink) {
			importLink = "https://classic.warcraftlogs.com/reports/HmXdtqRTwFchK89j#fight=31";
		}

		const url = new URL(importLink);
		var reportID = url.pathname.split("reports/")[1];
		var fightID = url.hash.split("=")[1];
		var settings = RaidSimSettings.create();
		var raid = Raid.create();
		raid.parties = new Array<Party>();
		settings.raid = raid;
		var encounter = Encounter.create();
		encounter.targets = new Array<Target>();
		var target = Target.create();
		// TODO: look up target in WCL data.
		target.armor = 7700;
		target.level = 73;
		target.mobType = MobType.MobTypeDemon;

		encounter.targets.push(target)
		settings.encounter = encounter;

		fetch("https://classic.warcraftlogs.com/oauth/token", {
			"method": "POST",
			"headers": {
				"Authorization": "Basic " + btoa("963d31c8-7efa-4dde-87cf-1b254a8a2f8c:lRJVhujEEnF96xfUoxVHSpnqKN9v8bTqGEjutsO3"),
			},
			body: new URLSearchParams({
				'grant_type': 'client_credentials'
			})
		}).then(response => response.json()).then(data => {
			fetch("https://classic.warcraftlogs.com/api/v2/client?=", {
				"method": "POST",
				"headers": {
					"Content-Type": "application/json",
					"Authorization": "Bearer " + data.access_token,
				},
				"body": `{"query":"{reportData { report(code: \\"${reportID}\\") { guild { name faction {id} } playerDetails(fightIDs: [${fightID}], endTime: 99999999)  events(fightIDs: [${fightID}], dataType:CombatantInfo, endTime: 99999999) {data}  fights(fightIDs: [31]) { startTime endTime }}}}"}`
			}).then(response => {
				console.log(response);
				return response.json()
			}).then(data => {
				var playerDetails = new Map<Number, any>();

				encounter.duration = (data.data.reportData.report.fights[0].endTime - data.data.reportData.report.fights[0].startTime) / 1000;

				(data.data.reportData.report.playerDetails.data.playerDetails.tanks as Array<any>).forEach(player => playerDetails.set(player.id, player) );
				(data.data.reportData.report.playerDetails.data.playerDetails.dps as Array<any>).forEach(player => playerDetails.set(player.id, player) );
				(data.data.reportData.report.playerDetails.data.playerDetails.healers as Array<any>).forEach(player => playerDetails.set(player.id, player) );

				var currentParty = 0;
				raid.parties.push(Party.create());
				(data.data.reportData.report.events.data as Array<any>).forEach(info => {
					if (currentParty == 5) {
						return;
					}

					var player = Player.create();
					var details = playerDetails.get(info.sourceID);
					var spec = specNames[details.specs[0].spec];
					if (spec == null || spec == undefined) {
						return;
					}

					const matchingPresets = playerPresets.filter(preset => preset.spec == spec);
					if (matchingPresets.length == 0) {
						return;
					}
					const matchingPreset = matchingPresets[0];

					var specFuncs = specTypeFunctions[spec];
					player = withSpecProto(spec, player, matchingPreset.rotation, specFuncs.talentsCreate(), matchingPreset.specOptions);
					player.talentsString = matchingPreset.talents;
					player.consumes = matchingPreset.consumes;

					player.name = details.name;
					player.class = nameToClass(details.type);
					player.equipment = EquipmentSpec.create();
					player.equipment.items = new Array<ItemSpec>();

					// If we couldn't find the type, skip
					if (!player.spec || player.spec.oneofKind == undefined) {
						// TODO: lookup buffbot types
						return;
					}
					const faction = data.data.reportData.report.guild.faction.id as Faction;
					player.race = matchingPreset.defaultFactionRaces[faction];

					(info.gear as Array<any>).forEach(gear => {
						var item = ItemSpec.create();
						item.id = gear.id;
						const dbEnchant = this.simUI.sim.getEnchantFlexible(gear.permanentEnchant);
						if (dbEnchant) {
							item.enchant = dbEnchant.id;
						} else {
							item.enchant = 0;
						}
						if (gear.gems) {
							item.gems = new Array<number>();
							(gear.gems as Array<any>).forEach(gemInfo => item.gems.push(gemInfo.id));	
						}
						player.equipment!.items.push(item);
					});
					
					raid.parties[currentParty].players.push(player);
					if (raid.parties[currentParty].players.length == 5) {
						currentParty++;
					}
					if (raid.parties.length <= currentParty) {
						if (currentParty == 5) {
							return;
						}
						raid.parties.push(Party.create());
					}
				});

				this.simUI.fromProto(TypedEvent.nextEventID(), settings);
				this.close();
			}).catch(err => {
				console.error(err);
			});	
		});
	}
}


// Maps WCL spec names to internal Spec enum.
const specNames: Record<string, Spec> = {
	'Balance': Spec.SpecBalanceDruid,
	'Elemental': Spec.SpecElementalShaman,
	'Enhancement': Spec.SpecEnhancementShaman,
  	'Feral': Spec.SpecFeralDruid,
	'Survival': Spec.SpecHunter,
	'BeastMastery': Spec.SpecHunter,
	'Arcane': Spec.SpecMage,
	'Fire': Spec.SpecMage,
	'Frost': Spec.SpecMage,
	'Assassination': Spec.SpecRogue,
	'Combat': Spec.SpecRogue,
	'Retribution': Spec.SpecRetributionPaladin,
	// 'Justicar': Spec.SpecProtectionPaladin,
	'Shadow': Spec.SpecShadowPriest,
	'Smite': Spec.SpecSmitePriest,
	'Destruction': Spec.SpecWarlock,
	'Affliction': Spec.SpecWarlock,
	'Demonology': Spec.SpecWarlock,
	'Arms': Spec.SpecWarrior,
	'Fury': Spec.SpecWarrior,
	'Protection': Spec.SpecProtectionWarrior,
};
