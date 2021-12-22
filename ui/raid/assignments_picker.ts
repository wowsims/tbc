import { Component } from '/tbc/core/components/component.js';
import { Input, InputConfig } from '/tbc/core/components/input.js';
import { RaidTargetPicker } from '/tbc/core/components/raid_target_picker.js';
import { Player } from '/tbc/core/player.js';
import { Raid } from '/tbc/core/raid.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { Class } from '/tbc/core/proto/common.js';
import { RaidTarget } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { getEnumValues } from '/tbc/core/utils.js';
import { wait } from '/tbc/core/utils.js';
import { newRaidTarget, emptyRaidTarget, NO_TARGET } from '/tbc/core/proto_utils/utils.js';

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

interface AssignmentTargetPicker {
	playerOrBot: Player<any> | BuffBot,
	targetPicker: RaidTargetPicker<Player<any> | BuffBot>,
	targetPlayer: Player<any> | null;
};

export class InnervatesPicker extends Component {
	readonly raidSimUI: RaidSimUI;
	readonly changeEmitter: TypedEvent<void> = new TypedEvent<void>();

	private readonly playersContainer: HTMLElement;

	private targetPickers: Array<AssignmentTargetPicker>;

  constructor(parentElem: HTMLElement, raidSimUI: RaidSimUI) {
    super(parentElem, 'innervates-picker-root');
		this.raidSimUI = raidSimUI;
		this.targetPickers = [];

		this.playersContainer = document.createElement('div');
		this.playersContainer.classList.add('innervate-players-container', 'settings-section');
		this.rootElem.appendChild(this.playersContainer);

		this.update(this.raidSimUI.getPlayersAndBuffBots());
		this.raidSimUI.compChangeEmitter.on(eventID => {
			this.recoverRaidTargets();
			this.update(this.raidSimUI.getPlayersAndBuffBots());
		});
	}

	private update(playersAndBots: Array<Player<any> | BuffBot | null>) {
		this.playersContainer.innerHTML = `
			<label>Innervates</label>
		`;

		const druids = playersAndBots.filter(playerOrBot => playerOrBot?.getClass() == Class.ClassDruid) as Array<Player<any> | BuffBot>;
		if (druids.length == 0) {
			this.rootElem.style.display = 'none';
		} else {
			this.rootElem.style.display = 'initial';
		}

		this.targetPickers = druids.map((druid, druidIndex) => {
			const row = document.createElement('div');
			row.classList.add('innervate-player');
			this.playersContainer.appendChild(row);

			const innervateSourceElem = RaidTargetPicker.makeOptionElem({
				iconUrl: druid instanceof Player ? druid.getTalentTreeIcon() : druid.settings.iconUrl,
				text: druid.getLabel(),
				color: druid.getClassColor(),
				isDropdown: false,
			});
			innervateSourceElem.classList.add('raid-target-picker-root');
			row.appendChild(innervateSourceElem);

			const arrow = document.createElement('span');
			arrow.classList.add('innervate-arrow', 'fa', 'fa-arrow-right');
			row.appendChild(arrow);

			let raidTargetPicker: RaidTargetPicker<Player<any> | BuffBot> | null = null;
			if (druid instanceof Player) {
				raidTargetPicker = new RaidTargetPicker<Player<any>>(row, druid, {
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
								isDropdown: true,
								value: newRaidTarget(player!.getRaidIndex()),
							};
						});
					},

					changedEvent: (player: Player<any>) => player.specOptionsChangeEmitter,
					getValue: (player: Player<any>) => (player.getSpecOptions() as DruidOptions).innervateTarget || emptyRaidTarget(),
					setValue: (eventID: EventID, player: Player<any>, newValue: RaidTarget) => {
						const newOptions = player.getSpecOptions() as DruidOptions;
						newOptions.innervateTarget = newValue;
						player.setSpecOptions(eventID, newOptions);
					},
				});
			} else {
				raidTargetPicker = new RaidTargetPicker<BuffBot>(row, druid, {
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
								isDropdown: true,
								value: newRaidTarget(player!.getRaidIndex()),
							};
						});
					},

					changedEvent: (player: BuffBot) => player.innervateAssignmentChangeEmitter,
					getValue: (player: BuffBot) => player.getInnervateAssignment(),
					setValue: (eventID: EventID, player: BuffBot, newValue: RaidTarget) => player.setInnervateAssignment(eventID, newValue),
				});
			}

			const targetPickerData = {
				playerOrBot: druid,
				targetPicker: raidTargetPicker!,
				targetPlayer: this.raidSimUI.sim.raid.getPlayerFromRaidTarget(raidTargetPicker!.getInputValue()),
			};

			raidTargetPicker!.changeEmitter.on(eventID => {
				targetPickerData.targetPlayer = this.raidSimUI.sim.raid.getPlayerFromRaidTarget(raidTargetPicker!.getInputValue());
			});

			return targetPickerData;
		});
	}

	// Tries to recover the current raid targets after the raid comp has changed.
	// For example, if an innervate is targeted onto a specific mage and that mage is
	// moved, we want to keep targeting the same mage.
	//
	// Note that when two characters are swapped, multiple compChange events are fired
	// and one of the characters will momentarily not be part of the raid. To address
	// this we have to wait a bit before checking.
	private recoverRaidTargets() {
		const oldTargetPickers = this.targetPickers.slice();
		const oldPlayerTargets = oldTargetPickers.map(otp => otp.targetPlayer);

		// TODO: This needs to somehow reference the 'parent' eventID so that undo
		// actions undo them together.
		const eventID = TypedEvent.nextEventID();
		TypedEvent.freezeAllAndDo(() => {
			oldTargetPickers.forEach((targetPicker, i) => {
				const oldPlayerOrBot = targetPicker.playerOrBot;
				const oldPlayerTarget = oldPlayerTargets[i];
				const newPlayersAndBots = this.raidSimUI.getPlayersAndBuffBots();

				if (!newPlayersAndBots.includes(oldPlayerOrBot))
					return;

				if (!oldPlayerTarget || !newPlayersAndBots.includes(oldPlayerTarget))
					return;

				const raidTarget = newRaidTarget(oldPlayerTarget.getRaidIndex());

				if (oldPlayerOrBot instanceof Player) {
					const newOptions = oldPlayerOrBot.getSpecOptions() as DruidOptions;
					newOptions.innervateTarget = raidTarget;
					oldPlayerOrBot.setSpecOptions(eventID, newOptions);
				} else {
					oldPlayerOrBot.setInnervateAssignment(eventID, raidTarget);
				}
			});
		});
	}
}
