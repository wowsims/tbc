import { Component } from '/tbc/core/components/component.js';
import { RaidTargetPicker } from '/tbc/core/components/raid_target_picker.js';
import { Player } from '/tbc/core/player.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { newRaidTarget, emptyRaidTarget } from '/tbc/core/proto_utils/utils.js';
import { BuffBot } from './buff_bot.js';
export class AssignmentsPicker extends Component {
    constructor(parentElem, raidSimUI) {
        super(parentElem, 'assignments-picker-root');
        this.changeEmitter = new TypedEvent();
        this.raidSimUI = raidSimUI;
        this.innervatesPicker = new InnervatesPicker(this.rootElem, raidSimUI);
        this.powerInfusionsPicker = new PowerInfusionsPicker(this.rootElem, raidSimUI);
    }
}
;
class AssignedBuffPicker extends Component {
    constructor(parentElem, raidSimUI) {
        super(parentElem, 'assigned-buff-picker-root');
        this.changeEmitter = new TypedEvent();
        this.raidSimUI = raidSimUI;
        this.targetPickers = [];
        this.playersContainer = document.createElement('fieldset');
        this.playersContainer.classList.add('assigned-buff-players-container', 'settings-section');
        this.rootElem.appendChild(this.playersContainer);
        this.update();
        this.raidSimUI.changeEmitter.on(eventID => {
            // Disabled because this is bugged.
            //this.recoverRaidTargets();
            this.update();
        });
    }
    update() {
        this.playersContainer.innerHTML = `
			<legend>${this.getTitle().toUpperCase()}</legend>
		`;
        const sourcePlayers = this.getSourcePlayers();
        if (sourcePlayers.length == 0) {
            this.rootElem.style.display = 'none';
        }
        else {
            this.rootElem.style.display = 'initial';
        }
        this.targetPickers = sourcePlayers.map((sourcePlayer, sourcePlayerIndex) => {
            const row = document.createElement('div');
            row.classList.add('assigned-buff-player');
            this.playersContainer.appendChild(row);
            const sourceElem = RaidTargetPicker.makeOptionElem({
                iconUrl: sourcePlayer instanceof Player ? sourcePlayer.getTalentTreeIcon() : sourcePlayer.settings.iconUrl,
                text: sourcePlayer.getLabel(),
                color: sourcePlayer.getClassColor(),
                isDropdown: false,
            });
            sourceElem.classList.add('raid-target-picker-root');
            row.appendChild(sourceElem);
            const arrow = document.createElement('span');
            arrow.classList.add('assigned-buff-arrow', 'fa', 'fa-arrow-right');
            row.appendChild(arrow);
            let raidTargetPicker = null;
            if (sourcePlayer instanceof Player) {
                raidTargetPicker = new RaidTargetPicker(row, sourcePlayer, {
                    extraCssClasses: [
                        'assigned-buff-target-picker',
                    ],
                    noTargetLabel: 'Unassigned',
                    compChangeEmitter: this.raidSimUI.sim.raid.compChangeEmitter,
                    getOptions: () => {
                        return this.raidSimUI.sim.raid.getPlayers().filter(player => player != null).map(player => {
                            return {
                                iconUrl: player.getTalentTreeIcon(),
                                text: player.getLabel(),
                                color: player.getClassColor(),
                                isDropdown: true,
                                value: newRaidTarget(player.getRaidIndex()),
                            };
                        });
                    },
                    changedEvent: (player) => player.specOptionsChangeEmitter,
                    getValue: (player) => this.getPlayerValue(player),
                    setValue: (eventID, player, newValue) => this.setPlayerValue(eventID, player, newValue),
                });
            }
            else {
                raidTargetPicker = new RaidTargetPicker(row, sourcePlayer, {
                    extraCssClasses: [
                        'assigned-buff-target-picker',
                    ],
                    noTargetLabel: 'Unassigned',
                    compChangeEmitter: this.raidSimUI.sim.raid.compChangeEmitter,
                    getOptions: () => {
                        return this.raidSimUI.sim.raid.getPlayers().filter(player => player != null).map(player => {
                            return {
                                iconUrl: player.getTalentTreeIcon(),
                                text: player.getLabel(),
                                color: player.getClassColor(),
                                isDropdown: true,
                                value: newRaidTarget(player.getRaidIndex()),
                            };
                        });
                    },
                    changedEvent: (buffBot) => buffBot.changeEmitter,
                    getValue: (buffBot) => this.getBuffBotValue(buffBot),
                    setValue: (eventID, buffBot, newValue) => this.setBuffBotValue(eventID, buffBot, newValue),
                });
            }
            const targetPickerData = {
                playerOrBot: sourcePlayer,
                targetPicker: raidTargetPicker,
                targetPlayer: this.raidSimUI.sim.raid.getPlayerFromRaidTarget(raidTargetPicker.getInputValue()),
            };
            raidTargetPicker.changeEmitter.on(eventID => {
                targetPickerData.targetPlayer = this.raidSimUI.sim.raid.getPlayerFromRaidTarget(raidTargetPicker.getInputValue());
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
    recoverRaidTargets() {
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
                //if (oldPlayerOrBot instanceof Player) {
                //	const newOptions = oldPlayerOrBot.getSpecOptions() as DruidOptions;
                //	newOptions.innervateTarget = raidTarget;
                //	oldPlayerOrBot.setSpecOptions(eventID, newOptions);
                //} else {
                //	oldPlayerOrBot.setInnervateAssignment(eventID, raidTarget);
                //}
            });
        });
    }
}
class InnervatesPicker extends AssignedBuffPicker {
    getTitle() {
        return 'Innervates';
    }
    getSourcePlayers() {
        return this.raidSimUI.getPlayersAndBuffBots().filter(playerOrBot => playerOrBot?.getClass() == Class.ClassDruid);
    }
    getPlayerValue(player) {
        return player.getSpecOptions().innervateTarget || emptyRaidTarget();
    }
    setPlayerValue(eventID, player, newValue) {
        const newOptions = player.getSpecOptions();
        newOptions.innervateTarget = newValue;
        player.setSpecOptions(eventID, newOptions);
    }
    getBuffBotValue(buffBot) {
        return buffBot.getInnervateAssignment();
    }
    setBuffBotValue(eventID, buffBot, newValue) {
        buffBot.setInnervateAssignment(eventID, newValue);
    }
}
class PowerInfusionsPicker extends AssignedBuffPicker {
    getTitle() {
        return 'Power Infusions';
    }
    getSourcePlayers() {
        return this.raidSimUI.getPlayersAndBuffBots()
            .filter(playerOrBot => playerOrBot?.getClass() == Class.ClassPriest)
            .filter(playerOrBot => {
            if (playerOrBot instanceof BuffBot) {
                return playerOrBot.settings.buffBotId == 'Divine Spirit Priest';
            }
            else {
                const player = playerOrBot;
                if (!player.getTalents().powerInfusion) {
                    return false;
                }
                // Don't include shadow priests even if they have the talent, because they
                // don't have a raid target option for this.
                return player.spec == Spec.SpecSmitePriest;
            }
        });
    }
    getPlayerValue(player) {
        return player.getSpecOptions().powerInfusionTarget || emptyRaidTarget();
    }
    setPlayerValue(eventID, player, newValue) {
        const newOptions = player.getSpecOptions();
        newOptions.powerInfusionTarget = newValue;
        player.setSpecOptions(eventID, newOptions);
    }
    getBuffBotValue(buffBot) {
        return buffBot.getPowerInfusionAssignment();
    }
    setBuffBotValue(eventID, buffBot, newValue) {
        buffBot.setPowerInfusionAssignment(eventID, newValue);
    }
}
