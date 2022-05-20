import { Exporter } from '/tbc/core/components/exporters.js';
import { Importer } from '/tbc/core/components/importers.js';
import { BuffBot, RaidSimSettings } from '/tbc/core/proto/ui.js';
import { Party, Player, Raid } from '../core/proto/api.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Class, Encounter, EquipmentSpec, ItemSpec, MobType, RaidTarget, Spec, Target } from '../core/proto/common.js';
import { nameToClass } from '../core/proto_utils/names.js';
import { Faction, makeDefaultBlessings, specTypeFunctions, withSpecProto } from '../core/proto_utils/utils.js';
import { playerPresets } from './presets.js';
export function newRaidImporters(simUI) {
    const importSettings = document.createElement('div');
    importSettings.classList.add('import-settings', 'sim-dropdown-menu');
    importSettings.innerHTML = `
		<span id="importMenuLink" class="dropdown-toggle fas fa-file-import" role="button" data-toggle="dropdown" aria-haspopup="true" arai-expanded="false"></span>
		<div class="dropdown-menu dropdown-menu-right" aria-labelledby="importMenuLink">
		</div>
	`;
    const linkElem = importSettings.getElementsByClassName('dropdown-toggle')[0];
    tippy(linkElem, {
        'content': 'Import',
        'allowHTML': true,
    });
    const menuElem = importSettings.getElementsByClassName('dropdown-menu')[0];
    const addMenuItem = (label, experimental, onClick) => {
        const itemElem = document.createElement('span');
        itemElem.classList.add('dropdown-item');
        itemElem.textContent = label;
        itemElem.addEventListener('click', onClick);
        menuElem.appendChild(itemElem);
        if (experimental) {
            itemElem.classList.add('experimental');
        }
    };
    addMenuItem('Json', false, () => new RaidJsonImporter(menuElem, simUI));
    addMenuItem('WCL', true, () => new RaidWCLImporter(menuElem, simUI));
    return importSettings;
}
export function newRaidExporters(simUI) {
    const exportSettings = document.createElement('div');
    exportSettings.classList.add('export-settings', 'sim-dropdown-menu');
    exportSettings.innerHTML = `
		<span id="exportMenuLink" class="dropdown-toggle fas fa-file-export" role="button" data-toggle="dropdown" aria-haspopup="true" arai-expanded="false"></span>
		<div class="dropdown-menu dropdown-menu-right" aria-labelledby="exportMenuLink">
		</div>
	`;
    const linkElem = exportSettings.getElementsByClassName('dropdown-toggle')[0];
    tippy(linkElem, {
        'content': 'Export',
        'allowHTML': true,
    });
    const menuElem = exportSettings.getElementsByClassName('dropdown-menu')[0];
    const addMenuItem = (label, onClick) => {
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
    constructor(parent, simUI) {
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
    onImport(data) {
        const settings = RaidSimSettings.fromJsonString(data);
        this.simUI.fromProto(TypedEvent.nextEventID(), settings);
        this.close();
    }
}
class RaidJsonExporter extends Exporter {
    constructor(parent, simUI) {
        super(parent, 'JSON Export', true);
        this.simUI = simUI;
        this.init();
    }
    getData() {
        return JSON.stringify(RaidSimSettings.toJson(this.simUI.toProto()), null, 2);
    }
}
class RaidWCLImporter extends Importer {
    constructor(parent, simUI) {
        super(parent, 'WCL Import');
        this.simUI = simUI;
        this.descriptionElem.innerHTML = `
			<p>
				WARNING: THIS IS EXPERIMENTAL
			</p>
			<p>
				Import entire raid from a WCL report.<br>
				The players will be out of order and any specs not implemented will not be imported.<br>
				If there are blank spots in the raid, just re-import.<br>
				Do not use the 'IMPORT FROM FILE' button, it won't work.<br>
			</p>
			<p>
				To import, paste the WCL report and fight link (https://classic.warcraftlogs.com/reports/REPORTID#fight=FIGHTID).<br>
				Include the fight ID or else first found fight will be used.<br>
			</p>
		`;
    }
    onImport(importLink) {
        // TODO: validate link so we dont get a crash
        const url = new URL(importLink);
        var reportID = url.pathname.split("reports/")[1];
        var hashVals = url.hash.replace("#", "").split("&");
        var fightID = "0";
        hashVals.forEach(val => {
            var parts = val.split("=");
            if (parts.length < 2 || parts[0] != "fight") {
                return;
            }
            fightID = parts[1];
        });
        var settings = RaidSimSettings.create();
        var raid = Raid.create();
        raid.parties = new Array();
        settings.raid = raid;
        var encounter = Encounter.create();
        encounter.targets = new Array();
        var target = Target.create();
        // TODO: look up target in WCL data.
        target.armor = 7700;
        target.level = 73;
        target.mobType = MobType.MobTypeDemon;
        var buffBots = new Array();
        encounter.targets.push(target);
        settings.encounter = encounter;
        var numPaladins = 0;
        var numPlayers = 0;
        // Raid index of players that recieved innervates
        var wclIDtoRaidIndex = new Map();
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
                "body": `{"query":"{reportData { report(code: \\"${reportID}\\") { guild { name faction {id} } playerDetails(fightIDs: [${fightID}], endTime: 99999999)  events(fightIDs: [${fightID}], dataType:CombatantInfo, endTime: 99999999) {data}  fights(fightIDs: [${fightID}]) { startTime endTime } buffs: events(fightIDs: [${fightID}], dataType:Buffs, endTime: 99999999, abilityID: 29166){ data }}}}"}`
            }).then(response => {
                console.log(response);
                return response.json();
            }).then(data => {
                var playerDetails = new Map();
                encounter.duration = (data.data.reportData.report.fights[0].endTime - data.data.reportData.report.fights[0].startTime) / 1000;
                data.data.reportData.report.playerDetails.data.playerDetails.tanks.forEach(player => playerDetails.set(player.id, player));
                data.data.reportData.report.playerDetails.data.playerDetails.dps.forEach(player => playerDetails.set(player.id, player));
                data.data.reportData.report.playerDetails.data.playerDetails.healers.forEach(player => playerDetails.set(player.id, player));
                var currentParty = 0;
                raid.parties.push(Party.create());
                data.data.reportData.report.events.data.forEach(info => {
                    if (currentParty == 5) {
                        return;
                    }
                    var player = Player.create();
                    var details = playerDetails.get(info.sourceID);
                    var spec = specNames[details.specs[0].spec];
                    if (details.type == "Paladin") {
                        numPaladins++;
                    }
                    // If we couldn't find the type, check buff bots.
                    if (spec == null || spec == undefined) {
                        var botID = buffBotNames[details.specs[0].spec + details.type];
                        if (botID == null || botID == undefined) {
                            console.log("Spec Not Implemented: ", details.specs[0].spec + details.type);
                            return;
                        }
                        // Insert bot!
                        var bot = BuffBot.create();
                        bot.id = botID;
                        bot.raidIndex = numPlayers;
                        buffBots.push(bot);
                        wclIDtoRaidIndex.set(info.sourceID, numPlayers);
                        numPlayers++;
                        // For now, insert a placeholder for the bot. It will get overwritten later I think.
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
                        return;
                    }
                    const matchingPresets = playerPresets.filter(preset => preset.spec == spec);
                    var presetIdx = 0;
                    if (matchingPresets.length == 0) {
                        return;
                    }
                    else if (matchingPresets.length > 1) {
                        var distance = 100;
                        // Search talents and find the preset that the players talents most closely match.
                        matchingPresets.forEach((preset, i) => {
                            var presetTalents = [0, 0, 0];
                            var talentIdx = 0;
                            // First sum up the number of talents per tree for preset.
                            Array.from(preset.talents).forEach((v) => {
                                if (v == '-') {
                                    talentIdx += 1;
                                    return;
                                }
                                presetTalents[talentIdx] += parseInt(v);
                            });
                            // Diff the distance to the preset.
                            var newDistance = Math.abs(info.talents[0].id - presetTalents[0]) + Math.abs(info.talents[1].id - presetTalents[1]) + Math.abs(info.talents[2].id - presetTalents[2]);
                            // If this is the best distance, assign this preset.
                            if (newDistance < distance) {
                                presetIdx = i;
                                distance = newDistance;
                            }
                        });
                    }
                    const matchingPreset = matchingPresets[presetIdx];
                    var specFuncs = specTypeFunctions[spec];
                    player = withSpecProto(spec, player, matchingPreset.rotation, specFuncs.talentsCreate(), matchingPreset.specOptions);
                    player.talentsString = matchingPreset.talents;
                    player.consumes = matchingPreset.consumes;
                    player.name = details.name;
                    player.class = nameToClass(details.type);
                    player.equipment = EquipmentSpec.create();
                    player.equipment.items = new Array();
                    // Default to UI setting
                    var faction = this.simUI.raidPicker?.getCurrentFaction();
                    // If defined in log, use that faction.
                    if (data.data.reportData.report.guild != null && data.data.reportData.report.guild != undefined) {
                        faction = data.data.reportData.report.guild.faction.id;
                    }
                    // Fallback if UI is broken and log has no faction.
                    if (faction == undefined) {
                        faction = Faction.Horde;
                    }
                    player.race = matchingPreset.defaultFactionRaces[faction];
                    info.gear.forEach(gear => {
                        var item = ItemSpec.create();
                        item.id = gear.id;
                        const dbEnchant = this.simUI.sim.getEnchantFlexible(gear.permanentEnchant);
                        if (dbEnchant) {
                            item.enchant = dbEnchant.id;
                        }
                        else {
                            item.enchant = 0;
                        }
                        if (gear.gems) {
                            item.gems = new Array();
                            gear.gems.forEach(gemInfo => item.gems.push(gemInfo.id));
                        }
                        player.equipment.items.push(item);
                    });
                    wclIDtoRaidIndex.set(info.sourceID, numPlayers);
                    numPlayers++;
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
                data.data.reportData.report.buffs.data.forEach(buff => {
                    if (buff.type == "removebuff") {
                        return;
                    }
                    // Innervate!
                    if (buff.abilityGameID == 29166) {
                        // Find if target is a player.
                        // If so, find source, could be a player or a buffbot.
                        // Apply innervate.
                        var sourceID = wclIDtoRaidIndex.get(buff.sourceID);
                        var targetID = wclIDtoRaidIndex.get(buff.targetID);
                        var source = settings.raid.parties[Math.floor(sourceID / 5)].players[sourceID % 5];
                        var target = settings.raid.parties[Math.floor(targetID / 5)].players[targetID % 5];
                        if (target.class == undefined) {
                            // Target is not a player
                            return;
                        }
                        if (source.class != Class.ClassDruid) {
                            // its a buffbot
                            buffBots.forEach(bot => {
                                if (bot.raidIndex == sourceID) {
                                    // Found our buffer
                                    bot.innervateAssignment = RaidTarget.create();
                                    bot.innervateAssignment.targetIndex = targetID;
                                    return;
                                }
                            });
                        }
                        else {
                            // Assign player sources
                            if (source.spec.oneofKind == "balanceDruid") {
                                source.spec.balanceDruid.options.innervateTarget = RaidTarget.create();
                                source.spec.balanceDruid.options.innervateTarget.targetIndex = targetID;
                            }
                            else if (source.spec.oneofKind == "feralDruid") {
                                source.spec.feralDruid.options.innervateTarget = RaidTarget.create();
                                source.spec.feralDruid.options.innervateTarget.targetIndex = targetID;
                            }
                            else if (source.spec.oneofKind == "feralTankDruid") {
                                source.spec.feralTankDruid.options.innervateTarget = RaidTarget.create();
                                source.spec.feralTankDruid.options.innervateTarget.targetIndex = targetID;
                            }
                        }
                    }
                });
                settings.blessings = makeDefaultBlessings(numPaladins);
                this.simUI.clearRaid(TypedEvent.nextEventID());
                this.simUI.fromProto(TypedEvent.nextEventID(), settings);
                this.simUI.setBuffBots(TypedEvent.nextEventID(), buffBots);
                this.close();
            }).catch(err => {
                console.error(err);
            });
        });
    }
}
// Maps WCL spec names to internal Spec enum.
const specNames = {
    'Balance': Spec.SpecBalanceDruid,
    'Elemental': Spec.SpecElementalShaman,
    'Enhancement': Spec.SpecEnhancementShaman,
    // 'Feral': Spec.SpecFeralDruid,
    // 'Warden': Spec.SpecFeralTankDruid,
    // 'Guardian': Spec.SpecFeralTankDruid,
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
    // 'Protection': Spec.SpecProtectionWarrior,
};
// Maps WCL spec+type to internal buff bot IDs.
const buffBotNames = {
    // DPS
    'FeralDruid': 'Bear',
    // Tank
    'JusticarPaladin': 'JoW Paladin',
    'ProtectionWarrior': 'Prot Warrior',
    'WardenDruid': 'Bear',
    'GuardianDruid': 'Bear',
    // Healers
    'HolyPaladin': 'Paladin',
    'HolyPriest': 'Holy Priest',
    'DisciplinePriest': 'Divine Spirit Priest',
    'RestorationDruid': 'Resto Druid',
    'RestorationShaman': 'Resto Shaman',
};
