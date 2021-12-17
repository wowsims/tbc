import { CloseButton } from '/tbc/core/components/close_button.js';
import { Component } from '/tbc/core/components/component.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { MAX_PARTY_SIZE } from '/tbc/core/party.js';
import { Player } from '/tbc/core/player.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { classColors } from '/tbc/core/proto_utils/utils.js';
import { specToClass } from '/tbc/core/proto_utils/utils.js';
import { newTalentsPicker } from '/tbc/core/talents/factory.js';
import { TypedEvent } from '/tbc/core/typed_event.js';
import { getEnumValues } from '/tbc/core/utils.js';
import { hexToRgba } from '/tbc/core/utils.js';
import { BalanceDruidSimUI } from '/tbc/balance_druid/sim.js';
import { ElementalShamanSimUI } from '/tbc/elemental_shaman/sim.js';
import { ShadowPriestSimUI } from '/tbc/shadow_priest/sim.js';
const NEW_PLAYER = -1;
var DragType;
(function (DragType) {
    DragType[DragType["New"] = 0] = "New";
    DragType[DragType["Move"] = 1] = "Move";
    DragType[DragType["Swap"] = 2] = "Swap";
    DragType[DragType["Copy"] = 3] = "Copy";
})(DragType || (DragType = {}));
;
export class RaidPicker extends Component {
    constructor(parent, raidSimUI, presets, buffBots) {
        super(parent, 'raid-picker-root');
        this.buffBotChangeEmitter = new TypedEvent();
        // Hold data about the player being dragged while the drag is happening.
        this.currentDragPlayer = null;
        this.currentDragPlayerFromIndex = NEW_PLAYER;
        this.currentDragType = DragType.New;
        this.raidSimUI = raidSimUI;
        this.raid = raidSimUI.sim.raid;
        this.presets = presets;
        this.buffBots = buffBots;
        const raidViewer = document.createElement('div');
        raidViewer.classList.add('current-raid-viewer');
        this.rootElem.appendChild(raidViewer);
        raidViewer.innerHTML = `
			<div class="parties-container">
			</div>
		`;
        const partiesContainer = this.rootElem.getElementsByClassName('parties-container')[0];
        this.partyPickers = this.raid.getParties().map((party, i) => new PartyPicker(partiesContainer, party, i, this));
        const newPlayerPickerRoot = document.createElement('div');
        newPlayerPickerRoot.classList.add('new-player-picker');
        this.rootElem.appendChild(newPlayerPickerRoot);
        this.newPlayerPicker = new NewPlayerPicker(newPlayerPickerRoot, this);
        this.rootElem.ondragend = event => {
            // Uncomment to remove player when dropped 'off' the raid.
            //if (this.currentDragPlayerFromIndex != NEW_PLAYER) {
            //	const playerPicker = this.getPlayerPicker(this.currentDragPlayerFromIndex);
            //	playerPicker.setPlayer(null);
            //}
            this.clearDragPlayer();
        };
    }
    getCurrentFaction() {
        return this.newPlayerPicker.currentFaction;
    }
    getCurrentPhase() {
        return this.raid.sim.getPhase();
    }
    getPlayerPicker(raidIndex) {
        return this.partyPickers[Math.floor(raidIndex / MAX_PARTY_SIZE)].playerPickers[raidIndex % MAX_PARTY_SIZE];
    }
    getPlayerPickers() {
        return [...new Array(25).keys()].map(i => this.getPlayerPicker(i));
    }
    getBuffBots() {
        return this.getPlayerPickers()
            .filter(picker => picker.player != null && 'buffBotId' in picker.player)
            .map(picker => {
            return {
                buffBot: picker.player,
                partyIndex: picker.partyPicker.index,
                raidIndex: picker.raidIndex,
            };
        });
    }
    setDragPlayer(player, fromIndex, type) {
        this.clearDragPlayer();
        this.currentDragPlayer = player;
        this.currentDragPlayerFromIndex = fromIndex;
        this.currentDragType = type;
        if (fromIndex != NEW_PLAYER) {
            const playerPicker = this.getPlayerPicker(fromIndex);
            playerPicker.rootElem.classList.add('dragFrom');
        }
    }
    clearDragPlayer() {
        if (this.currentDragPlayerFromIndex != NEW_PLAYER) {
            const playerPicker = this.getPlayerPicker(this.currentDragPlayerFromIndex);
            playerPicker.rootElem.classList.remove('dragFrom');
        }
        this.currentDragPlayer = null;
        this.currentDragPlayerFromIndex = NEW_PLAYER;
        this.currentDragType = DragType.New;
    }
}
export class PartyPicker extends Component {
    constructor(parent, party, index, raidPicker) {
        super(parent, 'party-picker-root');
        this.party = party;
        this.index = index;
        this.raidPicker = raidPicker;
        this.rootElem.innerHTML = `
			<div class="party-header">
				<span>Group ${index + 1}</span>
			</div>
			<div class="players-container">
			</div>
		`;
        const playersContainer = this.rootElem.getElementsByClassName('players-container')[0];
        this.playerPickers = [...Array(MAX_PARTY_SIZE).keys()].map(i => new PlayerPicker(playersContainer, this, i));
    }
}
export class PlayerPicker extends Component {
    constructor(parent, partyPicker, index) {
        super(parent, 'player-picker-root');
        this.index = index;
        this.raidIndex = partyPicker.index * MAX_PARTY_SIZE + index;
        this.player = null;
        this.partyPicker = partyPicker;
        this.raidPicker = partyPicker.raidPicker;
        this.rootElem.innerHTML = `
			<div class="player-label">
				<img class="player-icon"></img>
				<span class="player-name" contenteditable></span>
			</div>
			<div class="player-options">
				<span class="player-edit fa fa-edit"></span>
				<span class="player-copy fa fa-copy" draggable="true"></span>
				<span class="player-delete fa fa-times"></span>
			</div>
			<div class="player-results">
				<span class="player-results-dps"></span>
				<span class="player-results-reference-delta"></span>
			</div>
		`;
        this.labelElem = this.rootElem.getElementsByClassName('player-label')[0];
        this.iconElem = this.rootElem.getElementsByClassName('player-icon')[0];
        this.nameElem = this.rootElem.getElementsByClassName('player-name')[0];
        this.dpsResultElem = this.rootElem.getElementsByClassName('player-results-dps')[0];
        this.referenceDeltaElem = this.rootElem.getElementsByClassName('player-results-reference-delta')[0];
        this.nameElem.addEventListener('input', event => {
            if (this.player instanceof Player) {
                this.player.setName(this.nameElem.textContent || '');
            }
        });
        const maxLength = 25;
        this.nameElem.addEventListener('keydown', event => {
            // 9 is tab, 13 is enter
            if (event.keyCode == 9 || event.keyCode == 13) {
                event.preventDefault();
                const realPlayerPickers = this.raidPicker.getPlayerPickers().filter(pp => pp.player instanceof Player);
                const indexOfThis = realPlayerPickers.indexOf(this);
                if (indexOfThis != -1 && realPlayerPickers.length > indexOfThis + 1) {
                    realPlayerPickers[indexOfThis + 1].nameElem.focus();
                }
                else {
                    this.nameElem.blur();
                }
            }
            // escape
            if (event.keyCode == 27) {
                this.nameElem.blur();
            }
            // 8 is backspace, 46 is delete, 
            if ((event.keyCode != 8 && event.keyCode != 46) && (this.nameElem.textContent?.length || 0) >= maxLength) {
                event.preventDefault();
            }
        });
        const emptyName = 'Unnamed';
        this.nameElem.addEventListener('focusin', event => {
            const selection = window.getSelection();
            if (selection) {
                const range = document.createRange();
                range.selectNodeContents(this.nameElem);
                selection.removeAllRanges();
                selection.addRange(range);
            }
        });
        this.nameElem.addEventListener('focusout', event => {
            if (!this.nameElem.textContent) {
                this.nameElem.textContent = emptyName;
                if (this.player instanceof Player) {
                    this.player.setName(emptyName);
                }
            }
        });
        const dragStart = (event, type) => {
            if (this.player == null) {
                event.preventDefault();
                return;
            }
            const iconSrc = this.iconElem.src;
            const dragImage = new Image();
            dragImage.src = iconSrc;
            event.dataTransfer.setDragImage(dragImage, 30, 30);
            event.dataTransfer.setData("text/plain", iconSrc);
            event.dataTransfer.dropEffect = 'move';
            this.raidPicker.setDragPlayer(this.player, this.raidIndex, type);
        };
        this.labelElem.ondragstart = event => {
            dragStart(event, DragType.Swap);
        };
        const copyElem = this.rootElem.getElementsByClassName('player-copy')[0];
        tippy(copyElem, {
            'content': 'Copy',
            'allowHTML': true,
        });
        copyElem.ondragstart = event => {
            dragStart(event, DragType.Copy);
        };
        const deleteElem = this.rootElem.getElementsByClassName('player-delete')[0];
        tippy(deleteElem, {
            'content': 'Delete',
            'allowHTML': true,
        });
        deleteElem.addEventListener('click', event => {
            this.setPlayer(null);
        });
        let dragEnterCounter = 0;
        this.rootElem.ondragenter = event => {
            event.preventDefault();
            dragEnterCounter++;
            this.rootElem.classList.add('dragto');
        };
        this.rootElem.ondragleave = event => {
            event.preventDefault();
            dragEnterCounter--;
            if (dragEnterCounter <= 0) {
                this.rootElem.classList.remove('dragto');
            }
        };
        this.rootElem.ondragover = event => {
            event.preventDefault();
        };
        this.rootElem.ondrop = event => {
            event.preventDefault();
            dragEnterCounter = 0;
            this.rootElem.classList.remove('dragto');
            if (this.raidPicker.currentDragPlayer == null) {
                return;
            }
            if (this.raidPicker.currentDragPlayerFromIndex == this.raidIndex) {
                this.raidPicker.clearDragPlayer();
                return;
            }
            const dragType = this.raidPicker.currentDragType;
            if (this.raidPicker.currentDragPlayerFromIndex != NEW_PLAYER) {
                const fromPlayerPicker = this.raidPicker.getPlayerPicker(this.raidPicker.currentDragPlayerFromIndex);
                if (dragType == DragType.Swap) {
                    fromPlayerPicker.setPlayer(this.player);
                    fromPlayerPicker.iconElem.src = this.iconElem.src;
                }
                else if (dragType == DragType.Move) {
                    fromPlayerPicker.setPlayer(null);
                }
            }
            if (dragType == DragType.Copy) {
                if ('buffBotId' in this.raidPicker.currentDragPlayer) {
                    this.setPlayer(this.raidPicker.currentDragPlayer);
                }
                else {
                    this.setPlayer(this.raidPicker.currentDragPlayer.clone());
                }
            }
            else {
                this.setPlayer(this.raidPicker.currentDragPlayer);
            }
            this.iconElem.src = event.dataTransfer.getData('text/plain');
            this.raidPicker.clearDragPlayer();
        };
        const editElem = this.rootElem.getElementsByClassName('player-edit')[0];
        tippy(editElem, {
            'content': 'Edit',
            'allowHTML': true,
        });
        editElem.addEventListener('click', event => {
            if (this.player instanceof Player) {
                new PlayerEditorModal(this.player);
            }
        });
        this.raidPicker.raidSimUI.referenceChangeEmitter.on(() => {
            const currentData = this.raidPicker.raidSimUI.getCurrentData();
            const referenceData = this.raidPicker.raidSimUI.getReferenceData();
            const playerDps = currentData?.simResult.getDamageMetrics({ player: this.raidIndex }).avg || 0;
            const referenceDps = referenceData?.simResult.getDamageMetrics({ player: this.raidIndex }).avg || 0;
            if (playerDps == 0 && referenceDps == 0) {
                this.dpsResultElem.textContent = '';
                this.referenceDeltaElem.textContent = '';
                return;
            }
            this.dpsResultElem.textContent = playerDps.toFixed(1);
            if (!referenceData) {
                this.referenceDeltaElem.textContent = '';
                return;
            }
            const delta = playerDps - referenceDps;
            const deltaStr = delta.toFixed(1);
            if (delta >= 0) {
                this.referenceDeltaElem.textContent = '+' + deltaStr;
                this.referenceDeltaElem.classList.remove('negative');
                this.referenceDeltaElem.classList.add('positive');
            }
            else {
                this.referenceDeltaElem.textContent = '' + deltaStr;
                this.referenceDeltaElem.classList.remove('positive');
                this.referenceDeltaElem.classList.add('negative');
            }
        });
        this.update();
    }
    setPlayer(newPlayer) {
        if (newPlayer == this.player) {
            return;
        }
        this.dpsResultElem.textContent = '';
        this.referenceDeltaElem.textContent = '';
        const oldPlayerWasBuffBot = this.player != null && 'buffBotId' in this.player;
        this.player = newPlayer;
        if (newPlayer == null || newPlayer instanceof Player) {
            this.partyPicker.party.setPlayer(this.index, newPlayer);
            if (oldPlayerWasBuffBot) {
                this.raidPicker.buffBotChangeEmitter.emit();
            }
        }
        else {
            this.raidPicker.buffBotChangeEmitter.emit();
        }
        this.update();
    }
    update() {
        if (this.player == null) {
            this.rootElem.classList.add('empty');
            this.rootElem.classList.remove('buff-bot');
            this.rootElem.style.backgroundColor = 'black';
            this.labelElem.setAttribute('draggable', 'false');
            this.nameElem.textContent = '';
            this.nameElem.removeAttribute('contenteditable');
        }
        else if ('buffBotId' in this.player) {
            this.rootElem.classList.remove('empty');
            this.rootElem.classList.add('buff-bot');
            this.rootElem.style.backgroundColor = classColors[specToClass[this.player.spec]];
            this.labelElem.setAttribute('draggable', 'true');
            this.nameElem.textContent = this.player.name;
            this.nameElem.removeAttribute('contenteditable');
        }
        else {
            this.rootElem.classList.remove('empty');
            this.rootElem.classList.remove('buff-bot');
            this.rootElem.style.backgroundColor = classColors[specToClass[this.player.spec]];
            this.labelElem.setAttribute('draggable', 'true');
            this.nameElem.textContent = this.player.getName();
            this.nameElem.setAttribute('contenteditable', '');
        }
    }
}
class PlayerEditorModal extends Component {
    constructor(player) {
        super(document.body, 'player-editor-modal');
        this.rootElem.id = 'playerEditorModal';
        this.rootElem.innerHTML = `
			<div class="player-editor within-raid-sim">
			</div>
		`;
        new CloseButton(this.rootElem, () => {
            $('#playerEditorModal').bPopup().close();
            this.rootElem.remove();
        });
        const editorRoot = this.rootElem.getElementsByClassName('player-editor')[0];
        const individualSim = specSimFactories[player.spec](editorRoot, player);
        $('#playerEditorModal').bPopup({
            closeClass: 'player-editor-close',
            onClose: () => {
                this.rootElem.remove();
            },
        });
    }
}
class NewPlayerPicker extends Component {
    constructor(parent, raidPicker) {
        super(parent, 'new-player-picker-root');
        this.raidPicker = raidPicker;
        this.currentFaction = Faction.Alliance;
        this.rootElem.innerHTML = `
			<div class="new-player-picker-controls">
				<div class="faction-selector"></div>
				<div class="phase-selector"></div>
			</div>
			<div class="presets-container"></div>
			<div class="buff-bots-container">
				<div class="buff-bots-title">
					<span class="buff-bots-title-text">Buff Bots</span>
					<span class="buff-bots-tooltip fa fa-info-circle"></span>
				</div>
			</div>
		`;
        const factionSelector = new EnumPicker(this.rootElem.getElementsByClassName('faction-selector')[0], this, {
            label: 'Faction',
            labelTooltip: 'Default faction for newly-created players.',
            values: [
                { name: 'Alliance', value: Faction.Alliance },
                { name: 'Horde', value: Faction.Horde },
            ],
            changedEvent: (picker) => new TypedEvent(),
            getValue: (picker) => picker.currentFaction,
            setValue: (picker, newValue) => {
                picker.currentFaction = newValue;
            },
        });
        const phaseSelector = new EnumPicker(this.rootElem.getElementsByClassName('phase-selector')[0], this, {
            label: 'Phase',
            labelTooltip: 'Newly-created players will start with approximate BIS gear from this phase.',
            values: [
                { name: '1', value: 1 },
                { name: '2', value: 2 },
            ],
            changedEvent: (picker) => this.raidPicker.raid.sim.phaseChangeEmitter,
            getValue: (picker) => this.raidPicker.raid.sim.getPhase(),
            setValue: (picker, newValue) => {
                this.raidPicker.raid.sim.setPhase(newValue);
            },
        });
        const presetsContainer = this.rootElem.getElementsByClassName('presets-container')[0];
        getEnumValues(Class).forEach(wowClass => {
            if (wowClass == Class.ClassUnknown) {
                return;
            }
            const matchingPresets = this.raidPicker.presets.filter(preset => specToClass[preset.spec] == wowClass);
            if (matchingPresets.length == 0) {
                return;
            }
            const classPresetsContainer = document.createElement('div');
            classPresetsContainer.classList.add('class-presets-container');
            presetsContainer.appendChild(classPresetsContainer);
            classPresetsContainer.style.backgroundColor = hexToRgba(classColors[wowClass], 0.5);
            matchingPresets.forEach(matchingPreset => {
                const presetElem = document.createElement('div');
                presetElem.classList.add('preset-picker');
                classPresetsContainer.appendChild(presetElem);
                const presetIconElem = document.createElement('img');
                presetIconElem.classList.add('preset-picker-icon');
                presetElem.appendChild(presetIconElem);
                presetIconElem.src = matchingPreset.iconUrl;
                tippy(presetIconElem, {
                    'content': matchingPreset.tooltip,
                    'allowHTML': true,
                });
                presetElem.setAttribute('draggable', 'true');
                presetElem.ondragstart = event => {
                    const dragImage = new Image();
                    dragImage.src = matchingPreset.iconUrl;
                    event.dataTransfer.setDragImage(dragImage, 30, 30);
                    event.dataTransfer.setData("text/plain", matchingPreset.iconUrl);
                    event.dataTransfer.dropEffect = 'copy';
                    const newPlayer = new Player(matchingPreset.spec, this.raidPicker.raid.sim);
                    newPlayer.setRace(matchingPreset.defaultFactionRaces[this.raidPicker.getCurrentFaction()]);
                    newPlayer.setRotation(matchingPreset.rotation);
                    newPlayer.setTalentsString(matchingPreset.talents);
                    // TalentPicker needed to convert from talents string into talents proto.
                    newTalentsPicker(newPlayer.spec, document.createElement('div'), newPlayer);
                    newPlayer.setSpecOptions(matchingPreset.specOptions);
                    newPlayer.setConsumes(matchingPreset.consumes);
                    newPlayer.setName(matchingPreset.defaultName);
                    // Need to wait because the gear might not be loaded yet.
                    this.raidPicker.raid.sim.waitForInit().then(() => {
                        newPlayer.setGear(this.raidPicker.raid.sim.lookupEquipmentSpec(matchingPreset.defaultGear[this.raidPicker.getCurrentFaction()][this.raidPicker.getCurrentPhase()]));
                    });
                    this.raidPicker.setDragPlayer(newPlayer, NEW_PLAYER, DragType.New);
                };
            });
        });
        const buffbotsTooltip = this.rootElem.getElementsByClassName('buff-bots-tooltip')[0];
        tippy(buffbotsTooltip, {
            'content': 'Buff bots do not do DPS or any actions at all, except to buff their raid/party members. They are used as placeholders for classes we haven\'t implemented yet, or never will (e.g. healers) so that a proper raid environment can still be simulated.',
            'allowHTML': true,
        });
        const buffbotsContainer = this.rootElem.getElementsByClassName('buff-bots-container')[0];
        getEnumValues(Class).forEach(wowClass => {
            if (wowClass == Class.ClassUnknown) {
                return;
            }
            const matchingBuffBots = this.raidPicker.buffBots.filter(buffBot => specToClass[buffBot.spec] == wowClass);
            if (matchingBuffBots.length == 0) {
                return;
            }
            const classPresetsContainer = document.createElement('div');
            classPresetsContainer.classList.add('class-presets-container');
            buffbotsContainer.appendChild(classPresetsContainer);
            classPresetsContainer.style.backgroundColor = hexToRgba(classColors[wowClass], 0.5);
            matchingBuffBots.forEach(matchingBuffBot => {
                const presetElem = document.createElement('div');
                presetElem.classList.add('preset-picker');
                presetElem.classList.add('preset-picker-buff-bot');
                classPresetsContainer.appendChild(presetElem);
                const presetIconElem = document.createElement('img');
                presetIconElem.classList.add('preset-picker-icon');
                presetElem.appendChild(presetIconElem);
                presetIconElem.src = matchingBuffBot.iconUrl;
                tippy(presetIconElem, {
                    'content': matchingBuffBot.tooltip,
                    'allowHTML': true,
                });
                presetElem.setAttribute('draggable', 'true');
                presetElem.ondragstart = event => {
                    const dragImage = new Image();
                    dragImage.src = matchingBuffBot.iconUrl;
                    event.dataTransfer.setDragImage(dragImage, 30, 30);
                    event.dataTransfer.setData("text/plain", matchingBuffBot.iconUrl);
                    event.dataTransfer.dropEffect = 'copy';
                    this.raidPicker.setDragPlayer(matchingBuffBot, NEW_PLAYER, DragType.New);
                };
            });
        });
    }
}
export const specSimFactories = {
    [Spec.SpecBalanceDruid]: (parentElem, player) => new BalanceDruidSimUI(parentElem, player),
    [Spec.SpecElementalShaman]: (parentElem, player) => new ElementalShamanSimUI(parentElem, player),
    [Spec.SpecShadowPriest]: (parentElem, player) => new ShadowPriestSimUI(parentElem, player),
};
