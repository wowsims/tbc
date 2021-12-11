import { specToLocalStorageKey } from '/tbc/core/components/component.js';
import { SimOptions } from '/tbc/core/proto/api.js';
import { specToLocalStorageKey } from '/tbc/core/proto_utils/utils.js';

import { Raid } from './raid.js';
import { Sim } from './sim.js';
import { Encounter } from './encounter.js';
import { Target } from './target.js';
import { TypedEvent } from './typed_event.js';

declare var tippy: any;
declare var pako: any;

const CURRENT_SETTINGS_STORAGE_KEY = '__currentSettings__';

export interface SimUIConfig {
	knownIssues?: Array<string>;
}

// Core UI module.
export abstract class SimUI extends Component {
  readonly sim: Sim;
  readonly raid: Raid;
  readonly encounter: Encounter;

  // Emits when anything from the sim, raid, or encounter changes.
  readonly changeEmitter = new TypedEvent<void>();

  constructor(parentElem: HTMLElement, sim: Sim, config: SimUIConfig) {
		super(parentElem, 'sim-ui');
    this.sim = sim;
    this.rootElem.innerHTML = simHTML;

		this.raid = new Raid(this.sim);
    this.encounter = new Encounter(this.sim);

    [
      this.sim.changeEmitter,
      this.raid.changeEmitter,
      this.encounter.changeEmitter,
    ].forEach(emitter => emitter.on(() => this.changeEmitter.emit()));

		Array.from(document.getElementsByClassName('known-issues')).forEach(element => {
			if (config.knownIssues?.length) {
				(element as HTMLElement).style.display = 'initial';
			} else {
				return;
			}

			
			tippy(element, {
				'content': `
				<ul class="known-issues-tooltip">
					${config.knownIssues.map(issue => '<li>' + issue + '</li>').join('')}
				</ul>
				`,
				'allowHTML': true,
			});
		});
  }

  // Returns JSON representing all the current values.
  toJson(): Object {
    return {
      'raid': this.raid.toJson(),
      'encounter': this.encounter.toJson(),
    };
	}

  // Set all the current values, assumes obj is the same type returned by toJson().
  fromJson(obj: any) {
		// For legacy format. Do not remove this until 2022/01/05 (1 month).
		if (obj['sim']) {
			if (!obj['raid']) {
				obj['raid'] = {
					'parties': [
						{
							'players': [
								{
									'spec': this.player.spec,
									'player': obj['player'],
								},
							],
							'buffs': obj['sim']['partyBuffs'],
						},
					],
					'buffs': obj['sim']['raidBuffs'],
				};
				obj['raid']['parties'][0]['players'][0]['player']['buffs'] = obj['sim']['individualBuffs'];
			}
		}

		if (obj['raid']) {
			this.raid.fromJson(obj['raid']);
		}

		if (obj['encounter']) {
			this.encounter.fromJson(obj['encounter']);
		}
  }

  async init(): Promise<void> {
    await this.sim.init();

    this.changeEmitter.on(() => {
      const jsonStr = JSON.stringify(this.toJson());
      window.localStorage.setItem(this.getStorageKey(CURRENT_SETTINGS_STORAGE_KEY), jsonStr);
    });
  }

  makeRaidSimRequest(iterations: number, debug: boolean): RaidSimRequest {
		return RaidSimRequest.create({
			raid: this.raid.toProto(),
			encounter: this.encounter.toProto(),
			simOptions: SimOptions.create({
				iterations: iterations,
				debug: debug,
			}),
		});
  }

	abstract getStorageKey(postfix: string): string;
}

const simHTML = `
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
    </div>
  </section>
</div>
`;
