import { Encounter } from '/tbc/core/encounter.js';
import { Party } from '/tbc/core/party.js';
import { Player } from '/tbc/core/player.js';
import { Raid } from '/tbc/core/raid.js';
import { Sim } from '/tbc/core/sim.js';
import { Target } from '/tbc/core/target.js';
import { Actions } from '/tbc/core/components/actions.js';
import { BooleanPicker, BooleanPickerConfig } from '/tbc/core/components/boolean_picker.js';
import { CharacterStats } from '/tbc/core/components/character_stats.js';
import { BonusStatsPicker } from '/tbc/core/components/bonus_stats_picker.js';
import { DetailedResults } from '/tbc/core/components/detailed_results.js';
import { EncounterPicker, EncounterPickerConfig } from '/tbc/core/components/encounter_picker.js';
import { EnumPicker, EnumPickerConfig } from '/tbc/core/components/enum_picker.js';
import { GearPicker } from '/tbc/core/components/gear_picker.js';
import { IconInput } from '/tbc/core/components/icon_picker.js';
import { IconPicker } from '/tbc/core/components/icon_picker.js';
import { LogRunner } from '/tbc/core/components/log_runner.js';
import { NumberPicker, NumberPickerConfig } from '/tbc/core/components/number_picker.js';
import { Results } from '/tbc/core/components/results.js';
import { SavedDataConfig } from '/tbc/core/components/saved_data_manager.js';
import { SavedDataManager } from '/tbc/core/components/saved_data_manager.js';
import { RaidBuffs } from '/tbc/core/proto/common.js';
import { PartyBuffs } from '/tbc/core/proto/common.js';
import { IndividualBuffs } from '/tbc/core/proto/common.js';
import { Class } from '/tbc/core/proto/common.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { Encounter as EncounterProto } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Stat } from '/tbc/core/proto/common.js';
import { Gear } from '/tbc/core/proto_utils/gear.js';
import { raceNames } from '/tbc/core/proto_utils/names.js';
import { Stats } from '/tbc/core/proto_utils/stats.js';
import { SpecRotation } from '/tbc/core/proto_utils/utils.js';
import { specNames } from '/tbc/core/proto_utils/utils.js';
import { specToEligibleRaces } from '/tbc/core/proto_utils/utils.js';
import { newTalentsPicker } from '/tbc/core/talents/factory.js';
import { equalsOrBothNull } from '/tbc/core/utils.js';

import { SimUI, SimUIConfig } from '/tbc/core/sim_ui.js';

import * as Tooltips from '/tbc/core/constants/tooltips.js';

declare var Muuri: any;
declare var tippy: any;

export interface InputSection {
	tooltip?: string,
	inputs: Array<{
		type: 'boolean',
		cssClass: string,
		getModObject: (simUI: SimUI<any>) => any,
		config: BooleanPickerConfig<any>,
	} |
	{
		type: 'number',
		cssClass: string,
		getModObject: (simUI: SimUI<any>) => any,
		config: NumberPickerConfig<any>,
	} |
	{
		type: 'enum',
		cssClass: string,
		getModObject: (simUI: SimUI<any>) => any,
		config: EnumPickerConfig<any>,
	}>,
}

export interface DefaultThemeConfig<SpecType extends Spec> extends SimUIConfig<SpecType> {
  epStats: Array<Stat>;
  epReferenceStat: Stat;
  displayStats: Array<Stat>;

	selfBuffInputs: Array<IconInput<Player<any>>>,
	raidBuffInputs: Array<IconInput<Raid>>,
	partyBuffInputs: Array<IconInput<Party>>,
	playerBuffInputs: Array<IconInput<Player<any>>>,
	debuffInputs: Array<IconInput<Target>>;
	consumeInputs: Array<IconInput<Player<any>>>;
	rotationInputs: InputSection;
	otherInputs?: InputSection;
  additionalSections?: Record<string, InputSection>;
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

export class DefaultTheme<SpecType extends Spec> extends SimUI<SpecType> {
  private readonly _config: DefaultThemeConfig<SpecType>;

  constructor(parentElem: HTMLElement, config: DefaultThemeConfig<SpecType>) {
    super(parentElem, config)
    this._config = config;
    this.parentElem.innerHTML = layoutHTML;

		const titleElem = this.parentElem.getElementsByClassName('default-title')[0];
		titleElem.textContent = 'TBC ' + specNames[this.player.spec] + ' Sim';
		if (config.releaseStatus == 'Alpha') {
			titleElem.textContent += ' Alpha';
		} else if (config.releaseStatus == 'Beta') {
			titleElem.textContent += ' Beta';
		}

    const results = new Results(this.parentElem.getElementsByClassName('default-results')[0] as HTMLElement, this);
    const detailedResults = new DetailedResults(this.parentElem.getElementsByClassName('detailed-results')[0] as HTMLElement);
    const actions = new Actions(this.parentElem.getElementsByClassName('default-actions')[0] as HTMLElement, this, config.epStats, config.epReferenceStat, results, detailedResults);
    const logRunner = new LogRunner(this.parentElem.getElementsByClassName('log-runner')[0] as HTMLElement, this, results, detailedResults);

    const characterStats = new CharacterStats(this.parentElem.getElementsByClassName('default-stats')[0] as HTMLElement, config.displayStats, this.player);

    const gearPicker = new GearPicker(this.parentElem.getElementsByClassName('gear-picker')[0] as HTMLElement, this.player);
    const bonusStatsPicker = new BonusStatsPicker(this.parentElem.getElementsByClassName('bonus-stats-picker')[0] as HTMLElement, this.player, config.epStats);

    const talentsPicker = newTalentsPicker(this.player.spec, this.parentElem.getElementsByClassName('talents-picker')[0] as HTMLElement, this.player);
		// Add a url parameter to help people trapped in the wrong talents   ;)
		if (this._config.freezeTalents && !(new URLSearchParams(window.location.search).has('unlockTalents'))) {
			talentsPicker.freeze();
		}

    const settingsTab = document.getElementsByClassName('settings-inputs')[0] as HTMLElement;

		const configureIconSection = (sectionElem: HTMLElement, iconPickers: Array<IconPicker<any>>, tooltip?: string) => {
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

		const selfBuffsSection = this.parentElem.getElementsByClassName('self-buffs-section')[0] as HTMLElement;
    configureIconSection(
				selfBuffsSection,
				config.selfBuffInputs.map(iconInput => new IconPicker(selfBuffsSection, this.player, iconInput, this)),
				Tooltips.SELF_BUFFS_SECTION);

		const buffsSection = this.parentElem.getElementsByClassName('buffs-section')[0] as HTMLElement;
    configureIconSection(
				buffsSection,
				[
					config.raidBuffInputs.map(iconInput => new IconPicker(buffsSection, this.raid, iconInput, this)),
					config.playerBuffInputs.map(iconInput => new IconPicker(buffsSection, this.player, iconInput, this)),
					config.partyBuffInputs.map(iconInput => new IconPicker(buffsSection, this.party, iconInput, this)),
				].flat(),
				Tooltips.OTHER_BUFFS_SECTION);

		const debuffsSection = this.parentElem.getElementsByClassName('debuffs-section')[0] as HTMLElement;
    configureIconSection(
				debuffsSection,
				config.debuffInputs.map(iconInput => new IconPicker(debuffsSection, this.encounter.primaryTarget, iconInput, this)));

		const consumesSection = this.parentElem.getElementsByClassName('consumes-section')[0] as HTMLElement;
    configureIconSection(
				consumesSection,
				config.consumeInputs.map(iconInput => new IconPicker(consumesSection, this.player, iconInput, this)));

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
    configureInputSection(this.parentElem.getElementsByClassName('rotation-section')[0] as HTMLElement, config.rotationInputs);
		if (config.otherInputs?.inputs.length) {
      configureInputSection(this.parentElem.getElementsByClassName('other-settings-section')[0] as HTMLElement, config.otherInputs);
		}

		const makeInputSection = (sectionName: string, sectionConfig: InputSection) => {
			const sectionCssPrefix = sectionName.replace(/\s+/g, '');
      const sectionElem = document.createElement('section');
      sectionElem.classList.add('settings-section', sectionCssPrefix + '-section');
      sectionElem.innerHTML = `<label>${sectionName}</label>`;
      settingsTab.appendChild(sectionElem);
      configureInputSection(sectionElem, sectionConfig);
    };
		for (const [sectionName, sectionConfig] of Object.entries(config.additionalSections || {})) {
			makeInputSection(sectionName, sectionConfig);
    };

    const races = specToEligibleRaces[this.player.spec];
    const racePicker = new EnumPicker(this.parentElem.getElementsByClassName('race-section')[0] as HTMLElement, this.player, {
			values: races.map(race => {
				return {
					name: raceNames[race],
					value: race,
				};
			}),
      changedEvent: sim => sim.raceChangeEmitter,
      getValue: sim => sim.getRace(),
      setValue: (sim, newValue) => sim.setRace(newValue),
    });

    const encounterSectionElem = settingsTab.getElementsByClassName('encounter-section')[0] as HTMLElement;
		new EncounterPicker(encounterSectionElem, this.encounter, config.encounterPicker);

    // Init Muuri layout only when settings tab is clicked, because it needs the elements
    // to be shown so it can calculate sizes.
    let muuriInit = false;
    document.getElementById('settings-tab-toggle')!.addEventListener('click', event => {
      if (muuriInit) {
        return;
      }
      muuriInit = true;
      setTimeout(() => {
        new Muuri('.settings-inputs');
      }, 200); // Magic amount of time before Muuri init seems to work
    });
  }

  async init(): Promise<void> {
    const savedGearManager = new SavedDataManager<Player<any>, GearAndStats>(this.parentElem.getElementsByClassName('saved-gear-manager')[0] as HTMLElement, this.player, {
      label: 'Gear',
			storageKey: this.getSavedGearStorageKey(),
      getData: (player: Player<any>) => {
        return {
          gear: player.getGear(),
          bonusStats: player.getBonusStats(),
        };
      },
      setData: (player: Player<any>, newGearAndStats: GearAndStats) => {
        player.setGear(newGearAndStats.gear);
				if (newGearAndStats.bonusStats) {
					player.setBonusStats(newGearAndStats.bonusStats);
				}
      },
      changeEmitters: [this.player.changeEmitter],
      equals: (a: GearAndStats, b: GearAndStats) => a.gear.equals(b.gear) && equalsOrBothNull(a.bonusStats, b.bonusStats, (a, b) => a.equals(b)),
      toJson: (a: GearAndStats) => {
        return {
          gear: EquipmentSpec.toJson(a.gear.asSpec()),
          bonusStats: a.bonusStats?.toJson(),
        };
      },
      fromJson: (obj: any) => {
        return {
          gear: this.sim.lookupEquipmentSpec(EquipmentSpec.fromJson(obj['gear'])),
          bonusStats: Stats.fromJson(obj['bonusStats']),
        };
      },
    });

    const savedEncounterManager = new SavedDataManager<Encounter, EncounterProto>(this.parentElem.getElementsByClassName('saved-encounter-manager')[0] as HTMLElement, this.encounter, {
      label: 'Encounter',
			storageKey: this.getSavedEncounterStorageKey(),
      getData: (encounter: Encounter) => encounter.toProto(),
      setData: (encounter: Encounter, newEncounter: EncounterProto) => encounter.fromProto(newEncounter),
      changeEmitters: [this.encounter.changeEmitter],
      equals: (a: EncounterProto, b: EncounterProto) => EncounterProto.equals(a, b),
      toJson: (a: EncounterProto) => EncounterProto.toJson(a),
      fromJson: (obj: any) => EncounterProto.fromJson(obj),
    });

    const savedRotationManager = new SavedDataManager<Player<any>, SpecRotation<SpecType>>(this.parentElem.getElementsByClassName('saved-rotation-manager')[0] as HTMLElement, this.player, {
      label: 'Rotation',
			storageKey: this.getSavedRotationStorageKey(),
      getData: (player: Player<SpecType>) => player.getRotation(),
      setData: (player: Player<SpecType>, newRotation: SpecRotation<SpecType>) => player.setRotation(newRotation),
      changeEmitters: [this.player.rotationChangeEmitter],
      equals: (a: SpecRotation<SpecType>, b: SpecRotation<SpecType>) => this.player.specTypeFunctions.rotationEquals(a, b),
      toJson: (a: SpecRotation<SpecType>) => this.player.specTypeFunctions.rotationToJson(a),
      fromJson: (obj: any) => this.player.specTypeFunctions.rotationFromJson(obj),
    });

    const savedSettingsManager = new SavedDataManager<SimUI<any>, Settings>(this.parentElem.getElementsByClassName('saved-settings-manager')[0] as HTMLElement, this, {
      label: 'Settings',
			storageKey: this.getSavedSettingsStorageKey(),
      getData: (simUI: SimUI<any>) => {
        return {
          raidBuffs: simUI.raid.getBuffs(),
          partyBuffs: simUI.party.getBuffs(),
          individualBuffs: simUI.player.getBuffs(),
          consumes: simUI.player.getConsumes(),
          race: simUI.player.getRace(),
        };
      },
      setData: (simUI: SimUI<any>, newSettings: Settings) => {
        simUI.raid.setBuffs(newSettings.raidBuffs);
        simUI.party.setBuffs(newSettings.partyBuffs);
        simUI.player.setBuffs(newSettings.individualBuffs);
        simUI.player.setConsumes(newSettings.consumes);
        simUI.player.setRace(newSettings.race);
      },
      changeEmitters: [
				this.raid.buffsChangeEmitter,
				this.party.buffsChangeEmitter,
				this.player.buffsChangeEmitter,
				this.player.consumesChangeEmitter,
				this.player.raceChangeEmitter,
			],
      equals: (a: Settings, b: Settings) =>
					RaidBuffs.equals(a.raidBuffs, b.raidBuffs)
					&& PartyBuffs.equals(a.partyBuffs, b.partyBuffs)
					&& IndividualBuffs.equals(a.individualBuffs, b.individualBuffs)
					&& Consumes.equals(a.consumes, b.consumes)
					&& a.race == b.race,
      toJson: (a: Settings) => {
        return {
          raidBuffs: RaidBuffs.toJson(a.raidBuffs),
          partyBuffs: PartyBuffs.toJson(a.partyBuffs),
          individualBuffs: IndividualBuffs.toJson(a.individualBuffs),
          consumes: Consumes.toJson(a.consumes),
          race: a.race,
        };
      },
      fromJson: (obj: any) => {
        return {
          raidBuffs: RaidBuffs.fromJson(obj['raidBuffs']),
          partyBuffs: PartyBuffs.fromJson(obj['partyBuffs']),
          individualBuffs: IndividualBuffs.fromJson(obj['individualBuffs']),
          consumes: Consumes.fromJson(obj['consumes']),
          race: Number(obj['race']),
        };
      },
    });

    const savedTalentsManager = new SavedDataManager<Player<any>, string>(this.parentElem.getElementsByClassName('saved-talents-manager')[0] as HTMLElement, this.player, {
      label: 'Talents',
			storageKey: this.getSavedTalentsStorageKey(),
      getData: (player: Player<any>) => player.getTalentsString(),
      setData: (player: Player<any>, newTalentsString: string) => player.setTalentsString(newTalentsString),
      changeEmitters: [this.player.talentsStringChangeEmitter],
      equals: (a: string, b: string) => a == b,
      toJson: (a: string) => a,
      fromJson: (obj: any) => obj,
    });
		if (this._config.freezeTalents) {
			savedTalentsManager.freeze();
		}

    await super.init();

    savedGearManager.loadUserData();
    this._config.presets.gear.forEach(presetGear => {
      savedGearManager.addSavedData({
				name: presetGear.name,
				tooltip: presetGear.tooltip,
				isPreset: true,
				data: {
					gear: this.sim.lookupEquipmentSpec(presetGear.gear),
					bonusStats: new Stats(),
				},
				enableWhen: presetGear.enableWhen,
			});
    });

    savedEncounterManager.loadUserData();
    savedSettingsManager.loadUserData();

    savedTalentsManager.loadUserData();
    this._config.presets.talents.forEach(config => {
			config.isPreset = true;
      savedTalentsManager.addSavedData(config);
    });
  }
}

const layoutHTML = `
<div class="default-root">
  <section class="default-sidebar">
    <div class="default-title"></div>
    <div class="default-actions"></div>
    <div class="default-results"></div>
    <div class="default-stats"></div>
  </section>
  <section class="default-main">
    <ul class="nav nav-tabs">
      <li class="active"><a data-toggle="tab" href="#gear-tab">Gear</a></li>
      <li><a id="settings-tab-toggle" data-toggle="tab" href="#settings-tab">Settings</a></li>
      <li><a data-toggle="tab" href="#talents-tab">Talents</a></li>
      <li><a data-toggle="tab" href="#detailed-results-tab">Detailed Results</a></li>
      <li><a data-toggle="tab" href="#log-tab">Log</a></li>
      <li class="default-top-bar">
				<div class="known-issues">Known Issues</div>
				<span class="share-link fa fa-link"></span>
			</li>
    </ul>
    <div class="tab-content">
      <div id="gear-tab" class="tab-pane fade in active">
        <div class="gear-tab">
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
      </div>
      <div id="settings-tab" class="settings-tab tab-pane fade"">
        <div class="settings-inputs">
          <div class="settings-section-container">
            <section class="settings-section encounter-section">
              <label>Encounter</label>
            </section>
            <section class="settings-section race-section">
              <label>Race</label>
            </section>
            <section class="settings-section rotation-section">
              <label>Rotation</label>
            </section>
          </div>
          <div class="settings-section-container">
            <section class="settings-section self-buffs-section">
              <label>Self Buffs</label>
            </section>
          </div>
          <div class="settings-section-container">
            <section class="settings-section buffs-section">
              <label>Other Buffs</label>
            </section>
          </div>
          <div class="settings-section-container">
            <section class="settings-section consumes-section">
              <label>Consumes</label>
            </section>
          </div>
          <div class="settings-section-container">
            <section class="settings-section debuffs-section">
              <label>Debuffs</label>
            </section>
          </div>
          <div class="settings-section-container">
            <section class="settings-section other-settings-section">
              <label>Other</label>
            </section>
          </div>
        </div>
        <div class="settings-bottom-bar">
          <div class="saved-encounter-manager">
          </div>
          <div class="saved-rotation-manager">
          </div>
          <div class="saved-settings-manager">
          </div>
        </div>
      </div>
      <div id="talents-tab" class="tab-pane fade"">
        <div class="talents-tab">
          <div class="talents-picker">
          </div>
          <div class="saved-talents-manager">
          </div>
        </div>
      </div>
      <div id="detailed-results-tab" class="tab-pane fade">
				<div class="detailed-results">
				</div>
      </div>
      <div id="log-tab" class="tab-pane fade">
				<div class="log-runner">
				</div>
      </div>
    </div>
  </section>
</div>
`;
