import { Component } from '/tbc/core/components/component.js';
import { Input, InputConfig } from '/tbc/core/components/input.js';
import { RaidTargetPicker } from '/tbc/core/components/raid_target_picker.js';
import { Player } from '/tbc/core/player.js';
import { Raid } from '/tbc/core/raid.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Class } from '/tbc/core/proto/common.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { getEnumValues } from '/tbc/core/utils.js';
import { newRaidTarget, emptyRaidTarget } from '/tbc/core/proto_utils/utils.js';

import { BalanceDruid_Options as DruidOptions } from '/tbc/core/proto/druid.js';

import { BuffBot } from './buff_bot.js';
import { RaidSimUI } from './raid_sim_ui.js';

declare var tippy: any;

export class AssignmentsPicker extends Component {
	readonly raidSimUI: RaidSimUI;
	readonly changeEmitter: TypedEvent<void> = new TypedEvent<void>();

	private readonly innervatesPicker: InnervatesPicker;

  constructor(parentElem: HTMLElement, raidSimUI: RaidSimUI) {
    super(parentElem, 'assignments-picker-root');
		this.raidSimUI = raidSimUI;
		this.innervatesPicker = new InnervatesPicker(this.rootElem, raidSimUI);
	}
}

export class InnervatesPicker extends Component {
	readonly raidSimUI: RaidSimUI;
	readonly changeEmitter: TypedEvent<void> = new TypedEvent<void>();

	private readonly playersContainer: HTMLElement;

  constructor(parentElem: HTMLElement, raidSimUI: RaidSimUI) {
    super(parentElem, 'innervates-picker-root');
		this.raidSimUI = raidSimUI;

		this.playersContainer = document.createElement('div');
		this.playersContainer.classList.add('innervate-players-container', 'settings-section');
		this.rootElem.appendChild(this.playersContainer);

		this.update(this.raidSimUI.getPlayersAndBuffBots());
		this.raidSimUI.compChangeEmitter.on(() => {
			this.update(this.raidSimUI.getPlayersAndBuffBots());
		});
	}

	private update(playersAndBots: Array<Player<any> | BuffBot | null>) {
		this.playersContainer.innerHTML = `
			<label>Innervates</label>
		`;

		const druids = playersAndBots.filter(playerOrBot => playerOrBot?.getClass() == Class.ClassDruid) as Array<Player<any> | BuffBot>;

		druids.forEach(druid => {
			const row = document.createElement('div');
			row.classList.add('innervate-player');
			this.playersContainer.appendChild(row);

			if (druid instanceof Player) {
				const raidTargetPicker = new RaidTargetPicker<Player<any>>(row, druid, {
					extraCssClasses: [
						'innervate-target-picker',
					],
					noTargetLabel: 'Unassigned',
					compChangeEmitter: this.raidSimUI.sim.raid.compChangeEmitter,
					getOptions: () => {
						return this.raidSimUI.sim.raid.getPlayers().filter(player => player != null).map(player => {
							return {
								iconUrl: player!.getTalentTreeIcon(),
								text: player!.getLabel(),
								color: player!.getClassColor(),
								value: newRaidTarget(player!.getRaidIndex()),
							};
						});
					},
					changedEvent: (player: Player<any>) => player.specOptionsChangeEmitter,
					getValue: (player: Player<any>) => (player.getSpecOptions() as DruidOptions).innervateTarget || emptyRaidTarget(),
					setValue: (player: Player<any>, newValue: RaidTarget) => {
						const newOptions = player.getSpecOptions() as DruidOptions;
						newOptions.innervateTarget = newValue;
						player.setSpecOptions(newOptions);
					},
				});
			} else {
				const raidTargetPicker = new RaidTargetPicker<BuffBot>(row, druid, {
					extraCssClasses: [
						'innervate-target-picker',
					],
					noTargetLabel: 'Unassigned',
					compChangeEmitter: this.raidSimUI.sim.raid.compChangeEmitter,
					getOptions: () => {
						return this.raidSimUI.sim.raid.getPlayers().filter(player => player != null).map(player => {
							return {
								iconUrl: player!.getTalentTreeIcon(),
								text: player!.getLabel(),
								color: player!.getClassColor(),
								value: newRaidTarget(player!.getRaidIndex()),
							};
						});
					},
					changedEvent: (player: BuffBot) => player.innervateAssignmentChangeEmitter,
					getValue: (player: BuffBot) => player.getInnervateAssignment(),
					setValue: (player: BuffBot, newValue: RaidTarget) => player.setInnervateAssignment(newValue),
				});
			}
		});
	}
}
