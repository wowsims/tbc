import { BonusStatsPicker } from '/tbc/core/components/bonus_stats_picker.js';
import { BooleanPicker, BooleanPickerConfig } from '/tbc/core/components/boolean_picker.js';
import { CharacterStats } from '/tbc/core/components/character_stats.js';
import { Class } from '/tbc/core/proto/common.js';
import { CooldownsPicker } from '/tbc/core/components/cooldowns_picker.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Cooldowns } from '/tbc/core/proto/common.js';
import { Debuffs } from '/tbc/core/proto/common.js';
import { DetailedResults } from '/tbc/core/components/detailed_results.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { Encounter } from './encounter.js';
import { EncounterPicker, EncounterPickerConfig } from '/tbc/core/components/encounter_picker.js';
import { EnumPicker, EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { GearPicker } from '/tbc/core/components/gear_picker.js';
import { IconPicker, IconPickerConfig } from '/tbc/core/components/icon_picker.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { IndividualSimSettings } from '/tbc/core/proto/ui.js';
import { LogRunner } from '/tbc/core/components/log_runner.js';
import { NumberPicker, NumberPickerConfig } from '/tbc/core/components/number_picker.js';
import { Party } from './party.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { Player } from './player.js';
import { Player as PlayerProto } from '/tbc/core/proto/api.js';
import { Race } from '/tbc/core/proto/common.js';
import { Raid } from './raid.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { SavedDataConfig, SavedDataManager } from '/tbc/core/components/saved_data_manager.js';
import { SavedEncounter } from '/tbc/core/proto/ui.js';
import { SavedGearSet } from '/tbc/core/proto/ui.js';
import { SavedSettings } from '/tbc/core/proto/ui.js';
import { SavedTalents } from '/tbc/core/proto/ui.js';
import { Sim } from './sim.js';
import { SimOptions } from '/tbc/core/proto/api.js';
import { SimUI } from './sim_ui.js';
import { Spec } from '/tbc/core/proto/common.js';
import { getMetaGemConditionDescription } from '/tbc/core/proto_utils/gems.js';
import { SpecOptions } from '/tbc/core/proto_utils/utils.js';
import { SpecRotation } from '/tbc/core/proto_utils/utils.js';
import { Stat } from '/tbc/core/proto/common.js';
import { StatWeightsRequest } from '/tbc/core/proto/api.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { Target } from './target.js';
import { EventID, TypedEvent } from './typed_event.js';
import { addRaidSimAction, RaidSimResultsManager } from '/tbc/core/components/raid_sim_action.js';
import { addStatWeightsAction } from '/tbc/core/components/stat_weights_action.js';
import { equalsOrBothNull } from '/tbc/core/utils.js';
import { newTalentsPicker } from '/tbc/core/talents/factory.js';
import { raceNames } from '/tbc/core/proto_utils/names.js';
import { specNames } from '/tbc/core/proto_utils/utils.js';
import { specToEligibleRaces } from '/tbc/core/proto_utils/utils.js';
import { specToLocalStorageKey } from '/tbc/core/proto_utils/utils.js';

import * as Tooltips from '/tbc/core/constants/tooltips.js';

declare var Muuri: any;
declare var tippy: any;
declare var pako: any;

const SAVED_GEAR_STORAGE_KEY = '__savedGear__';
const SAVED_ROTATION_STORAGE_KEY = '__savedRotation__';
const SAVED_SETTINGS_STORAGE_KEY = '__savedSettings__';
const SAVED_TALENTS_STORAGE_KEY = '__savedTalents__';

export interface IndividualSimIconPickerConfig<ModObject, ValueType> extends IconPickerConfig<ModObject, ValueType> {
  // If set, all effects with matching tags will be deactivated when this
  // effect is enabled.
  exclusivityTags?: Array<ExclusivityTag>;
};

class IndividualSimIconPicker<ModObject, ValueType> extends IconPicker<ModObject, ValueType> {
  constructor(parent: HTMLElement, modObj: ModObject, input: IndividualSimIconPickerConfig<ModObject, ValueType>, simUI: IndividualSimUI<any>) {
		super(parent, modObj, input);

    if (input.exclusivityTags) {
      simUI.registerExclusiveEffect({
        tags: input.exclusivityTags,
        changedEvent: this.changeEmitter,
        isActive: () => Boolean(this.getInputValue()),
        deactivate: (eventID: EventID) => this.setValue(eventID, (typeof this.getInputValue() == 'number') ? 0 as unknown as ValueType : false as unknown as ValueType),
      });
    }
	}
}

export interface InputSection {
	tooltip?: string,
	inputs: Array<{
		type: 'boolean',
		getModObject: (simUI: IndividualSimUI<any>) => any,
		config: BooleanPickerConfig<any>,
	} |
	{
		type: 'number',
		getModObject: (simUI: IndividualSimUI<any>) => any,
		config: NumberPickerConfig<any>,
	} |
	{
		type: 'enum',
		getModObject: (simUI: IndividualSimUI<any>) => any,
		config: EnumPickerConfig<any>,
	}>,
}

export interface IndividualSimUIConfig<SpecType extends Spec> {
	// Additional css class to add to the root element.
	cssClass: string,

	knownIssues?: Array<string>;

  epStats: Array<Stat>;
  epReferenceStat: Stat;
  displayStats: Array<Stat>;
	modifyDisplayStats?: (player: Player<SpecType>, stats: Stats) => Stats,

  defaults: {
		gear: EquipmentSpec,
		epWeights: Stats,
    consumes: Consumes,
    rotation: SpecRotation<SpecType>,
    talents: string,
    specOptions: SpecOptions<SpecType>,

    raidBuffs: RaidBuffs,
    partyBuffs: PartyBuffs,
    individualBuffs: IndividualBuffs,

    debuffs: Debuffs,
  },

	selfBuffInputs: Array<IndividualSimIconPickerConfig<Player<any>, any>>,
	raidBuffInputs: Array<IndividualSimIconPickerConfig<Raid, any>>,
	partyBuffInputs: Array<IndividualSimIconPickerConfig<Party, any>>,
	playerBuffInputs: Array<IndividualSimIconPickerConfig<Player<any>, any>>,
	debuffInputs: Array<IndividualSimIconPickerConfig<Target, any>>;
	consumeInputs: Array<IndividualSimIconPickerConfig<Player<any>, any>>;
	rotationInputs: InputSection;
	otherInputs?: InputSection;

	// Extra UI sections with the same input config as other sections.
  additionalSections?: Record<string, InputSection>;

	// For when extra sections are needed, with even more flexibility than additionalSections.
	// Return value is the label for the section.
  customSections?: Array<(simUI: IndividualSimUI<SpecType>, parentElem: HTMLElement) => string>;

	encounterPicker: EncounterPickerConfig,
	freezeTalents?: boolean;

  presets: {
    gear: Array<PresetGear>,
    talents: Array<SavedDataConfig<Player<any>, string>>,
  },
}

export interface GearAndStats {
  gear: Gear,
  bonusStats?: Stats,
}

export interface PresetGear {
  name: string;
  gear: EquipmentSpec;
  tooltip?: string;
	enableWhen?: (obj: Player<any>) => boolean;
}

export interface Settings {
  raidBuffs: RaidBuffs,
  partyBuffs: PartyBuffs,
  individualBuffs: IndividualBuffs,
  consumes: Consumes,
  race: Race,
}

// Extended shared UI for all individual player sims.
export abstract class IndividualSimUI<SpecType extends Spec> extends SimUI {
  readonly player: Player<SpecType>;
	readonly individualConfig: IndividualSimUIConfig<SpecType>;

  private readonly exclusivityMap: Record<ExclusivityTag, Array<ExclusiveEffect>>;

	private raidSimResultsManager: RaidSimResultsManager | null;

	private settingsMuuri: any;

  constructor(parentElem: HTMLElement, player: Player<SpecType>, config: IndividualSimUIConfig<SpecType>) {
		super(parentElem, player.sim, {
			spec: player.spec,
			knownIssues: config.knownIssues,
		});
		this.rootElem.classList.add('individual-sim-ui', config.cssClass);
		this.player = player;
		this.individualConfig = config;
		this.raidSimResultsManager = null;
		this.settingsMuuri = null;
		this.addWarning({
			updateOn: this.player.gearChangeEmitter,
			shouldDisplay: () => this.player.getGear().hasInactiveMetaGem(),
			getContent: () => {
				const metaGem = this.player.getGear().getMetaGem()!;
				return `Meta gem disabled (${metaGem.name}): ${getMetaGemConditionDescription(metaGem)}`;
			},
		});

    this.exclusivityMap = {
      'Battle Elixir': [],
      'Drums': [],
      'Food': [],
      'Alchohol': [],
      'Guardian Elixir': [],
      'Potion': [],
      'Conjured': [],
      'Spirit': [],
      'Weapon Imbue': [],
    };

		if (!this.isWithinRaidSim) {
			// This needs to go before all the UI components so that gear loading is the
			// first callback invoked from waitForInit().
			this.sim.waitForInit().then(() => this.loadSettings());
		}
		this.player.setEpWeights(this.individualConfig.defaults.epWeights);

		this.addSidebarComponents();
		this.addTopbarComponents();
		this.addGearTab();
		this.addSettingsTab();
		this.addTalentsTab();

		if (!this.isWithinRaidSim) {
			this.addDetailedResultsTab();
			this.addLogTab();
		}
  }

	private loadSettings() {
		const initEventID = TypedEvent.nextEventID();
		TypedEvent.freezeAllAndDo(() => {
			let loadedSettings = false;

			let hash = window.location.hash;
			if (hash.length > 1) {
				// Remove leading '#'
				hash = hash.substring(1);
				try {
					let jsonData;
					if (new URLSearchParams(window.location.search).has('uncompressed')) {
						const jsonStr = atob(hash);
						const jsonData = JSON.parse(jsonStr);
						this.sim.fromJson(initEventID, jsonData, this.player.spec);
						loadedSettings = true;
					} else {
						const binary = atob(hash);
						const bytes = new Uint8Array(binary.length);
						for (let i = 0; i < bytes.length; i++) {
								bytes[i] = binary.charCodeAt(i);
						}
						const jsonStr = pako.inflate(bytes, { to: 'string' });  
						try {
							jsonData = JSON.parse(jsonStr);
						} catch (e) {
							// Json parse failure just means we're using the new proto format so its ok.
						}

						if (jsonData) {
							this.sim.fromJson(initEventID, jsonData, this.player.spec);
							loadedSettings = true;
						} else {
							// TODO: Deprecate the json pathways and make this the default on January 25th, 2022 (1 month).
							const settingsBytes = pako.inflate(bytes);  
							const settings = IndividualSimSettings.fromBinary(settingsBytes);
							this.fromProto(initEventID, settings);
						}
					}
				} catch (e) {
					console.warn('Failed to parse settings from window hash: ' + e);
				}
			}
			window.location.hash = '';

			const savedSettings = window.localStorage.getItem(this.getSettingsStorageKey());
			if (!loadedSettings && savedSettings != null) {
				try {
					const jsonData = JSON.parse(savedSettings);
					if (jsonData?.raid) {
						this.sim.fromJson(initEventID, jsonData, this.player.spec);
					} else {
						const settings = IndividualSimSettings.fromJsonString(savedSettings);
						this.fromProto(initEventID, settings);
					}
					loadedSettings = true;
				} catch (e) {
					console.warn('Failed to parse saved settings: ' + e);
				}
			}

			if (!loadedSettings) {
				this.applyDefaults(initEventID);
			}
			this.player.setName(initEventID, 'Player');

			// This needs to go last so it doesn't re-store things as they are initialized.
			this.changeEmitter.on(eventID => {
				const jsonStr = IndividualSimSettings.toJsonString(this.toProto());
				window.localStorage.setItem(this.getSettingsStorageKey(), jsonStr);
			});
		});
	}

	private addSidebarComponents() {
		this.raidSimResultsManager = addRaidSimAction(this);
		addStatWeightsAction(this, this.individualConfig.epStats, this.individualConfig.epReferenceStat);

    const characterStats = new CharacterStats(
				this.rootElem.getElementsByClassName('sim-sidebar-footer')[0] as HTMLElement,
				this.player,
				this.individualConfig.displayStats,
				this.individualConfig.modifyDisplayStats);
	}

	private addTopbarComponents() {
		const shareLink = document.createElement('span');
		shareLink.classList.add('share-link', 'fa', 'fa-link', 'within-raid-sim-hide');
		tippy(shareLink, {
			'content': 'Shareable link',
			'allowHTML': true,
		});
		shareLink.addEventListener('click', event => {
			const protoBytes = IndividualSimSettings.toBinary(this.toProto());
			const deflated = pako.deflate(protoBytes, { to: 'string' });
			const encoded = btoa(String.fromCharCode(...deflated));

			const linkUrl = new URL(window.location.href);
			linkUrl.hash = encoded;
			
			if (navigator.clipboard == undefined) {
				alert(linkUrl.toString());
			} else {
				navigator.clipboard.writeText(linkUrl.toString());
				alert('Current settings copied to clipboard!');
			}
		});
		this.addToolbarItem(shareLink);
	}

	private addGearTab() {
		this.addTab('GEAR', 'gear-tab', `
			<div class="gear-tab-columns">
				<div class="left-gear-panel">
					<div class="gear-picker">
					</div>
				</div>
				<div class="right-gear-panel">
					<div class="bonus-stats-picker">
					</div>
					<div class="saved-gear-manager">
					</div>
				</div>
			</div>
		`);

    const gearPicker = new GearPicker(this.rootElem.getElementsByClassName('gear-picker')[0] as HTMLElement, this.player);
    const bonusStatsPicker = new BonusStatsPicker(this.rootElem.getElementsByClassName('bonus-stats-picker')[0] as HTMLElement, this.player, this.individualConfig.epStats);

		const savedGearManager = new SavedDataManager<Player<any>, SavedGearSet>(this.rootElem.getElementsByClassName('saved-gear-manager')[0] as HTMLElement, this.player, {
			label: 'Gear',
			storageKey: this.getSavedGearStorageKey(),
			getData: (player: Player<any>) => {
				return SavedGearSet.create({
					gear: player.getGear().asSpec(),
					bonusStats: player.getBonusStats().asArray(),
				});
			},
			setData: (eventID: EventID, player: Player<any>, newSavedGear: SavedGearSet) => {
				TypedEvent.freezeAllAndDo(() => {
					player.setGear(eventID, this.sim.lookupEquipmentSpec(newSavedGear.gear || EquipmentSpec.create()));
					player.setBonusStats(eventID, new Stats(newSavedGear.bonusStats || []));
				});
			},
			changeEmitters: [this.player.changeEmitter],
			equals: (a: SavedGearSet, b: SavedGearSet) => SavedGearSet.equals(a, b),
			toJson: (a: SavedGearSet) => SavedGearSet.toJson(a),
			fromJson: (obj: any) => SavedGearSet.fromJson(obj),
		});

		this.sim.waitForInit().then(() => {
			savedGearManager.loadUserData();
			this.individualConfig.presets.gear.forEach(presetGear => {
				savedGearManager.addSavedData({
					name: presetGear.name,
					tooltip: presetGear.tooltip,
					isPreset: true,
					data: SavedGearSet.create({
						// Convert to gear and back so order is always the same.
						gear: this.sim.lookupEquipmentSpec(presetGear.gear).asSpec(),
						bonusStats: new Stats().asArray(),
					}),
					enableWhen: presetGear.enableWhen,
				});
			});
		});
	}

	private addSettingsTab() {
		this.addTab('SETTINGS', 'settings-tab', `
			<div class="settings-inputs">
				<div class="settings-section-container">
					<fieldset class="settings-section encounter-section within-raid-sim-hide">
						<legend>Encounter</legend>
					</fieldset>
					<fieldset class="settings-section race-section">
						<legend>Race</legend>
					</fieldset>
					<fieldset class="settings-section rotation-section">
						<legend>Rotation</legend>
					</fieldset>
				</div>
				<div class="settings-section-container custom-sections-container">
				</div>
				<div class="settings-section-container">
					<fieldset class="settings-section self-buffs-section">
						<legend>Self Buffs</legend>
					</fieldset>
				</div>
				<div class="settings-section-container within-raid-sim-hide">
					<fieldset class="settings-section buffs-section">
						<legend>Other Buffs</legend>
					</fieldset>
				</div>
				<div class="settings-section-container">
					<fieldset class="settings-section consumes-section">
						<legend>Consumes</legend>
					</fieldset>
				</div>
				<div class="settings-section-container cooldowns-section-container">
					<fieldset class="settings-section cooldowns-section">
						<legend>Cooldowns</legend>
						<div class="cooldowns-section-content">
						</div>
					</fieldset>
				</div>
				<div class="settings-section-container within-raid-sim-hide">
					<fieldset class="settings-section debuffs-section">
						<legend>Debuffs</legend>
					</fieldset>
				</div>
				<div class="settings-section-container">
					<fieldset class="settings-section other-settings-section">
						<legend>Other</legend>
					</fieldset>
				</div>
			</div>
			<div class="settings-bottom-bar">
				<div class="saved-encounter-manager within-raid-sim-hide">
				</div>
				<div class="saved-settings-manager within-raid-sim-hide">
				</div>
			</div>
		`);

    const settingsTab = this.rootElem.getElementsByClassName('settings-inputs')[0] as HTMLElement;

		const configureIconSection = (sectionElem: HTMLElement, iconPickers: Array<IconPicker<any, any>>, tooltip?: string) => {
			if (tooltip) {
				tippy(sectionElem, {
					'content': tooltip,
					'allowHTML': true,
				});
			}

			if (iconPickers.length == 0) {
				sectionElem.style.display = 'none';
			}
		};

		const selfBuffsSection = this.rootElem.getElementsByClassName('self-buffs-section')[0] as HTMLElement;
    configureIconSection(
				selfBuffsSection,
				this.individualConfig.selfBuffInputs.map(iconInput => new IndividualSimIconPicker(selfBuffsSection, this.player, iconInput, this)),
				Tooltips.SELF_BUFFS_SECTION);

		const buffsSection = this.rootElem.getElementsByClassName('buffs-section')[0] as HTMLElement;
    configureIconSection(
				buffsSection,
				[
					this.individualConfig.raidBuffInputs.map(iconInput => new IndividualSimIconPicker(buffsSection, this.sim.raid, iconInput, this)),
					this.individualConfig.playerBuffInputs.map(iconInput => new IndividualSimIconPicker(buffsSection, this.player, iconInput, this)),
					this.individualConfig.partyBuffInputs.map(iconInput => new IndividualSimIconPicker(buffsSection, this.player.getParty()!, iconInput, this)),
				].flat(),
				Tooltips.OTHER_BUFFS_SECTION);

		const debuffsSection = this.rootElem.getElementsByClassName('debuffs-section')[0] as HTMLElement;
    configureIconSection(
				debuffsSection,
				this.individualConfig.debuffInputs.map(iconInput => new IndividualSimIconPicker(debuffsSection, this.sim.encounter.primaryTarget, iconInput, this)),
				Tooltips.DEBUFFS_SECTION);

		const consumesSection = this.rootElem.getElementsByClassName('consumes-section')[0] as HTMLElement;
    configureIconSection(
				consumesSection,
				this.individualConfig.consumeInputs.map(iconInput => new IndividualSimIconPicker(consumesSection, this.player, iconInput, this)));

		const configureInputSection = (sectionElem: HTMLElement, sectionConfig: InputSection) => {
			if (sectionConfig.tooltip) {
				tippy(sectionElem, {
					'content': sectionConfig.tooltip,
					'allowHTML': true,
				});
			}

      sectionConfig.inputs.forEach(inputConfig => {
        if (inputConfig.type == 'number') {
          const picker = new NumberPicker(sectionElem, inputConfig.getModObject(this), inputConfig.config);
        } else if (inputConfig.type == 'boolean') {
          const picker = new BooleanPicker(sectionElem, inputConfig.getModObject(this), inputConfig.config);
        } else if (inputConfig.type == 'enum') {
          const picker = new EnumPicker(sectionElem, inputConfig.getModObject(this), inputConfig.config);
        }
      });
		};
    configureInputSection(this.rootElem.getElementsByClassName('rotation-section')[0] as HTMLElement, this.individualConfig.rotationInputs);
		if (this.individualConfig.otherInputs?.inputs.length) {
      configureInputSection(this.rootElem.getElementsByClassName('other-settings-section')[0] as HTMLElement, this.individualConfig.otherInputs);
		}

    const races = specToEligibleRaces[this.player.spec];
    const racePicker = new EnumPicker(this.rootElem.getElementsByClassName('race-section')[0] as HTMLElement, this.player, {
			values: races.map(race => {
				return {
					name: raceNames[race],
					value: race,
				};
			}),
      changedEvent: sim => sim.raceChangeEmitter,
      getValue: sim => sim.getRace(),
      setValue: (eventID, sim, newValue) => sim.setRace(eventID, newValue),
    });

    const encounterSectionElem = settingsTab.getElementsByClassName('encounter-section')[0] as HTMLElement;
		new EncounterPicker(encounterSectionElem, this.sim.encounter, this.individualConfig.encounterPicker);
    const savedEncounterManager = new SavedDataManager<Encounter, SavedEncounter>(this.rootElem.getElementsByClassName('saved-encounter-manager')[0] as HTMLElement, this.sim.encounter, {
      label: 'Encounter',
			storageKey: this.getSavedEncounterStorageKey(),
      getData: (encounter: Encounter) => SavedEncounter.create({ encounter: encounter.toProto() }),
      setData: (eventID: EventID, encounter: Encounter, newEncounter: SavedEncounter) => encounter.fromProto(eventID, newEncounter.encounter!),
      changeEmitters: [this.sim.encounter.changeEmitter],
      equals: (a: SavedEncounter, b: SavedEncounter) => SavedEncounter.equals(a, b),
      toJson: (a: SavedEncounter) => SavedEncounter.toJson(a),
      fromJson: (obj: any) => SavedEncounter.fromJson(obj),
    });

    const cooldownSectionElem = settingsTab.getElementsByClassName('cooldowns-section')[0] as HTMLElement;
    const cooldownContentElem = settingsTab.getElementsByClassName('cooldowns-section-content')[0] as HTMLElement;
		new CooldownsPicker(cooldownContentElem, this.player);
		tippy(cooldownSectionElem, {
			content: Tooltips.COOLDOWNS_SECTION,
			allowHTML: true,
			placement: 'left',
		});
		this.player.cooldownsChangeEmitter.on(() => {
			this.recomputeSettingsLayout();
		});

    // Init Muuri layout only when settings tab is clicked, because it needs the elements
    // to be shown so it can calculate sizes.
    (this.rootElem.getElementsByClassName('settings-tab-tab')[0] as HTMLElement)!.addEventListener('click', event => {
      if (this.settingsMuuri == null) {
				setTimeout(() => {
					this.settingsMuuri = new Muuri('.settings-inputs');
				}, 200); // Magic amount of time before Muuri init seems to work
      }

			setTimeout(() => {
				this.recomputeSettingsLayout();
			}, 200);
    });

    const savedSettingsManager = new SavedDataManager<IndividualSimUI<any>, SavedSettings>(this.rootElem.getElementsByClassName('saved-settings-manager')[0] as HTMLElement, this, {
      label: 'Settings',
			storageKey: this.getSavedSettingsStorageKey(),
      getData: (simUI: IndividualSimUI<any>) => {
        return SavedSettings.create({
          raidBuffs: simUI.sim.raid.getBuffs(),
          partyBuffs: simUI.player.getParty()?.getBuffs() || PartyBuffs.create(),
          playerBuffs: simUI.player.getBuffs(),
          consumes: simUI.player.getConsumes(),
          race: simUI.player.getRace(),
          cooldowns: simUI.player.getCooldowns(),
        });
      },
      setData: (eventID: EventID, simUI: IndividualSimUI<any>, newSettings: SavedSettings) => {
				TypedEvent.freezeAllAndDo(() => {
					simUI.sim.raid.setBuffs(eventID, newSettings.raidBuffs || RaidBuffs.create());
					const party = simUI.player.getParty();
					if (party) {
						party.setBuffs(eventID, newSettings.partyBuffs || PartyBuffs.create());
					}
					simUI.player.setBuffs(eventID, newSettings.playerBuffs || IndividualBuffs.create());
					simUI.player.setConsumes(eventID, newSettings.consumes || Consumes.create());
					simUI.player.setRace(eventID, newSettings.race);
					simUI.player.setCooldowns(eventID, newSettings.cooldowns || Cooldowns.create());
				});
      },
      changeEmitters: [
				this.sim.raid.buffsChangeEmitter,
				this.player.getParty()!.buffsChangeEmitter,
				this.player.buffsChangeEmitter,
				this.player.consumesChangeEmitter,
				this.player.raceChangeEmitter,
				this.player.cooldownsChangeEmitter,
			],
      equals: (a: SavedSettings, b: SavedSettings) => SavedSettings.equals(a, b),
      toJson: (a: SavedSettings) => SavedSettings.toJson(a),
      fromJson: (obj: any) => SavedSettings.fromJson(obj),
    });

    const customSectionsContainer = this.rootElem.getElementsByClassName('custom-sections-container')[0] as HTMLElement;
		let anyCustomSections = false;
		for (const [sectionName, sectionConfig] of Object.entries(this.individualConfig.additionalSections || {})) {
			const sectionCssPrefix = sectionName.replace(/\s+/g, '');
      const sectionElem = document.createElement('fieldset');
      sectionElem.classList.add('settings-section', sectionCssPrefix + '-section');
      sectionElem.innerHTML = `<legend>${sectionName}</legend>`;
      customSectionsContainer.appendChild(sectionElem);
      configureInputSection(sectionElem, sectionConfig);
			anyCustomSections = true;
    };

		(this.individualConfig.customSections || []).forEach(customSection => {
      const sectionElem = document.createElement('fieldset');
      customSectionsContainer.appendChild(sectionElem);
			const sectionName = customSection(this, sectionElem);
			const sectionCssPrefix = sectionName.replace(/\s+/g, '');
      sectionElem.classList.add('settings-section', sectionCssPrefix + '-section');
			const labelElem = document.createElement('legend');
			labelElem.textContent = sectionName;
			sectionElem.prepend(labelElem);
			anyCustomSections = true;
		});

		if (!anyCustomSections) {
			customSectionsContainer.remove();
		}

		this.sim.waitForInit().then(() => {
			savedEncounterManager.loadUserData();
			savedSettingsManager.loadUserData();
		});
	}

	private addTalentsTab() {
		this.addTab('TALENTS', 'talents-tab', `
			<div class="talents-picker">
			</div>
			<div class="saved-talents-manager">
			</div>
		`);

    const talentsPicker = newTalentsPicker(this.rootElem.getElementsByClassName('talents-picker')[0] as HTMLElement, this.player);
		const savedTalentsManager = new SavedDataManager<Player<any>, SavedTalents>(this.rootElem.getElementsByClassName('saved-talents-manager')[0] as HTMLElement, this.player, {
			label: 'Talents',
			storageKey: this.getSavedTalentsStorageKey(),
			getData: (player: Player<any>) => SavedTalents.create({
				talentsString: player.getTalentsString(),
			}),
			setData: (eventID: EventID, player: Player<any>, newTalents: SavedTalents) => player.setTalentsString(eventID, newTalents.talentsString),
			changeEmitters: [this.player.talentsStringChangeEmitter],
			equals: (a: SavedTalents, b: SavedTalents) => SavedTalents.equals(a, b),
			toJson: (a: SavedTalents) => SavedTalents.toJson(a),
			fromJson: (obj: any) => SavedTalents.fromJson(obj),
		});

		// Add a url parameter to help people trapped in the wrong talents   ;)
		const freezeTalents = this.individualConfig.freezeTalents && !(new URLSearchParams(window.location.search).has('unlockTalents'));
		if (freezeTalents) {
			savedTalentsManager.freeze();
			talentsPicker.freeze();
		}

		this.sim.waitForInit().then(() => {
			savedTalentsManager.loadUserData();
			this.individualConfig.presets.talents.forEach(config => {
				config.isPreset = true;
				savedTalentsManager.addSavedData({
					name: config.name,
					data: SavedTalents.create({
						talentsString: config.data,
					}),
				});
			});
		});
	}

	private addDetailedResultsTab() {
		this.addTab('DETAILED RESULTS', 'detailed-results-tab', `
			<div class="detailed-results">
			</div>
		`);

    const detailedResults = new DetailedResults(this.rootElem.getElementsByClassName('detailed-results')[0] as HTMLElement, this, this.raidSimResultsManager!);
	}

	private addLogTab() {
		this.addTab('LOG', 'log-tab', `
			<div class="log-runner">
			</div>
		`);

    const logRunner = new LogRunner(this.rootElem.getElementsByClassName('log-runner')[0] as HTMLElement, this);
	}

	private applyDefaults(eventID: EventID) {
		this.player.setGear(eventID, this.sim.lookupEquipmentSpec(this.individualConfig.defaults.gear));
		this.player.setConsumes(eventID, this.individualConfig.defaults.consumes);
		this.player.setRotation(eventID, this.individualConfig.defaults.rotation);
		this.player.setTalentsString(eventID, this.individualConfig.defaults.talents);
		this.player.setSpecOptions(eventID, this.individualConfig.defaults.specOptions);
		this.player.setBuffs(eventID, this.individualConfig.defaults.individualBuffs);
		this.player.getParty()!.setBuffs(eventID, this.individualConfig.defaults.partyBuffs);
		this.player.getRaid()!.setBuffs(eventID, this.individualConfig.defaults.raidBuffs);
		this.sim.encounter.primaryTarget.setDebuffs(eventID, this.individualConfig.defaults.debuffs);
	}

  registerExclusiveEffect(effect: ExclusiveEffect) {
    effect.tags.forEach(tag => {
      this.exclusivityMap[tag].push(effect);

      effect.changedEvent.on(eventID => {
        if (!effect.isActive())
          return;

				// TODO: Mark the parent somehow so we can track this for undo/redo.
				const newEventID = TypedEvent.nextEventID();
				TypedEvent.freezeAllAndDo(() => {
					this.exclusivityMap[tag].forEach(otherEffect => {
						if (otherEffect == effect || !otherEffect.isActive())
							return;

						otherEffect.deactivate(newEventID);
					});
				});
      });
    });
  }

	getSavedGearStorageKey(): string {
		return this.getStorageKey(SAVED_GEAR_STORAGE_KEY);
	}

	getSavedRotationStorageKey(): string {
		return this.getStorageKey(SAVED_ROTATION_STORAGE_KEY);
	}

	getSavedSettingsStorageKey(): string {
		return this.getStorageKey(SAVED_SETTINGS_STORAGE_KEY);
	}

	getSavedTalentsStorageKey(): string {
		return this.getStorageKey(SAVED_TALENTS_STORAGE_KEY);
	}

	recomputeSettingsLayout() {
		if (this.settingsMuuri) {
			//this.settingsMuuri.refreshItems();
		}
		window.dispatchEvent(new Event('resize'));
	}

	// Returns the actual key to use for local storage, based on the given key part and the site context.
	getStorageKey(keyPart: string): string {
		// Local storage is shared by all sites under the same domain, so we need to use
		// different keys for each spec site.
		return specToLocalStorageKey[this.player.spec] + keyPart;
	}

	toProto(): IndividualSimSettings {
		return IndividualSimSettings.create({
			player: this.player.toProto(),
			raidBuffs: this.sim.raid.getBuffs(),
			partyBuffs: this.player.getParty()?.getBuffs() || PartyBuffs.create(),
			encounter: this.sim.encounter.toProto(),
		});
	}

	fromProto(eventID: EventID, settings: IndividualSimSettings) {
		TypedEvent.freezeAllAndDo(() => {
			if (!settings.player) {
				return;
			}
			this.player.fromProto(eventID, settings.player);
			this.sim.raid.setBuffs(eventID, settings.raidBuffs || RaidBuffs.create());
			const party = this.player.getParty();
			if (party) {
				party.setBuffs(eventID, settings.partyBuffs || PartyBuffs.create());
			}
			this.sim.encounter.fromProto(eventID, settings.encounter || EncounterProto.create());
		});
	}
}

export type ExclusivityTag =
    'Battle Elixir'
    | 'Drums'
    | 'Food'
    | 'Alchohol'
    | 'Guardian Elixir'
    | 'Potion'
    | 'Conjured'
    | 'Spirit'
    | 'Weapon Imbue';

export interface ExclusiveEffect {
  tags: Array<ExclusivityTag>;
  changedEvent: TypedEvent<any>;
  isActive: () => boolean;
  deactivate: (eventID: EventID) => void;
}
