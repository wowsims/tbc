import { Component } from '/tbc/core/components/component.js';
import { IconEnumPicker } from '/tbc/core/components/icon_enum_picker.js';
import { NumberListPicker } from '/tbc/core/components/number_list_picker.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { ActionID as ActionIdProto } from '/tbc/core/proto/common.js';
import { Cooldown } from '/tbc/core/proto/common.js';
import { protoToActionId } from '/tbc/core/proto_utils/action_id.js';
import { getFullActionName } from '/tbc/core/resources.js';
export class CooldownsPicker extends Component {
    constructor(parentElem, player) {
        super(parentElem, 'cooldowns-picker-root');
        this.player = player;
        this.cooldownPickers = [];
        TypedEvent.onAny([this.player.currentStatsEmitter]).on(eventID => {
            this.update();
        });
    }
    update() {
        this.rootElem.innerHTML = '';
        const cooldowns = this.player.getCooldowns().cooldowns;
        this.cooldownPickers = [];
        for (let i = 0; i < cooldowns.length + 1; i++) {
            const cooldown = cooldowns[i];
            const row = document.createElement('div');
            row.classList.add('cooldown-picker');
            if (i == cooldowns.length) {
                row.classList.add('add-cooldown-picker');
            }
            this.rootElem.appendChild(row);
            const deleteButton = document.createElement('span');
            deleteButton.classList.add('delete-cooldown', 'fa', 'fa-times');
            deleteButton.addEventListener('click', event => {
                const newCooldowns = this.player.getCooldowns();
                newCooldowns.cooldowns.splice(i, 1);
                this.player.setCooldowns(TypedEvent.nextEventID(), newCooldowns);
            });
            row.appendChild(deleteButton);
            const actionPicker = this.makeActionPicker(row, i);
            const label = document.createElement('span');
            label.classList.add('cooldown-picker-label');
            if (cooldown && cooldown.id) {
                getFullActionName(protoToActionId(cooldown.id), this.player.getRaidIndex()).then(actionName => label.textContent = actionName);
            }
            row.appendChild(label);
            const timingsPicker = this.makeTimingsPicker(row, i);
            this.cooldownPickers.push(row);
        }
    }
    makeActionPicker(parentElem, cooldownIndex) {
        const availableCooldowns = this.player.getCurrentStats().cooldowns;
        const actionPicker = new IconEnumPicker(parentElem, this.player, {
            extraCssClasses: [
                'cooldown-action-picker',
            ],
            numColumns: 3,
            values: [
                { color: '#grey', value: ActionIdProto.create() },
            ].concat(availableCooldowns.map(cooldownAction => {
                return { actionId: protoToActionId(cooldownAction), value: cooldownAction };
            })),
            equals: (a, b) => ActionIdProto.equals(a, b),
            zeroValue: ActionIdProto.create(),
            backupIconUrl: (value) => protoToActionId(value),
            changedEvent: (player) => player.cooldownsChangeEmitter,
            getValue: (player) => player.getCooldowns().cooldowns[cooldownIndex]?.id || ActionIdProto.create(),
            setValue: (eventID, player, newValue) => {
                const newCooldowns = player.getCooldowns();
                while (newCooldowns.cooldowns.length < cooldownIndex) {
                    newCooldowns.cooldowns.push(Cooldown.create());
                }
                newCooldowns.cooldowns[cooldownIndex] = Cooldown.create({
                    id: newValue,
                    timings: [],
                });
                player.setCooldowns(eventID, newCooldowns);
            },
        });
        return actionPicker;
    }
    makeTimingsPicker(parentElem, cooldownIndex) {
        const actionPicker = new NumberListPicker(parentElem, this.player, {
            extraCssClasses: [
                'cooldown-timings-picker',
            ],
            placeholder: '20, 40, ...',
            changedEvent: (player) => player.cooldownsChangeEmitter,
            getValue: (player) => {
                return player.getCooldowns().cooldowns[cooldownIndex]?.timings || [];
            },
            setValue: (eventID, player, newValue) => {
                const newCooldowns = player.getCooldowns();
                newCooldowns.cooldowns[cooldownIndex].timings = newValue;
                player.setCooldowns(eventID, newCooldowns);
            },
            enableWhen: (player) => {
                const curCooldown = player.getCooldowns().cooldowns[cooldownIndex];
                return curCooldown && !ActionIdProto.equals(curCooldown.id, ActionIdProto.create());
            },
        });
        return actionPicker;
    }
}
