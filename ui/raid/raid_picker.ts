import { Component } from '/tbc/core/components/component.js';
import { EnumPicker } from '/tbc/core/components/enum_picker.js';
import { makePhaseSelector } from '/tbc/core/components/other_inputs.js';
import { Raid } from '/tbc/core/raid.js';
import { MAX_PARTY_SIZE } from '/tbc/core/party.js';
import { Party } from '/tbc/core/party.js';
import { Player } from '/tbc/core/player.js';
import { Consumes } from '/tbc/core/proto/common.js';
import { EquipmentSpec } from '/tbc/core/proto/common.js';
import { Race } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Faction } from '/tbc/core/proto_utils/utils.js';
import { SpecOptions } from '/tbc/core/proto_utils/utils.js';
import { SpecRotation } from '/tbc/core/proto_utils/utils.js';
import { TypedEvent } from '/tbc/core/typed_event.js';


export class RaidPicker extends Component {
	private readonly raid: Raid;
  private readonly partyPickers: Array<PartyPicker>;

  constructor(parent: HTMLElement, raid: Raid, specs: Array<PresetSpecSettings<any>>) {
    super(parent, 'raid-picker-root');
		this.raid = raid;

    const raidViewer = document.createElement('div');
    raidViewer.classList.add('current-raid-viewer');
    this.rootElem.appendChild(raidViewer);
		raidViewer.innerHTML = `
			<div class="parties-container">
			</div>
		`;

    const partiesContainer = this.rootElem.getElementsByClassName('parties-container')[0] as HTMLDivElement;
		this.partyPickers = this.raid.getParties().map((party, i) => new PartyPicker(partiesContainer, party, i));

    const newPlayerPickerRoot = document.createElement('div');
    newPlayerPickerRoot.classList.add('new-player-picker');
    this.rootElem.appendChild(newPlayerPickerRoot);

		const newPlayerPicker = new NewPlayerPicker(newPlayerPickerRoot, this.raid, specs);
	}
}

export class PartyPicker extends Component {
	private readonly party: Party;
	private readonly index: number;
  private readonly playerPickers: Array<PlayerPicker>;

  constructor(parent: HTMLElement, party: Party, index: number) {
    super(parent, 'party-picker-root');
		this.party = party;
		this.index = index;

		this.rootElem.innerHTML = `
			<div class="party-header">
				<span>Group ${index + 1}</span>
			</div>
			<div class="players-container">
			</div>
		`;

    const playersContainer = this.rootElem.getElementsByClassName('players-container')[0] as HTMLDivElement;
		this.playerPickers = [...Array(MAX_PARTY_SIZE).keys()].map(i => new PlayerPicker(playersContainer, this.party, i));
	}
}

export class PlayerPicker extends Component {
	private readonly party: Party;
	private readonly playerIndex: number;
	private player: Player<any> | null;

	private readonly nameElem: HTMLSpanElement;

  constructor(parent: HTMLElement, party: Party, playerIndex: number) {
    super(parent, 'player-picker-root');
		this.party = party;
		this.playerIndex = playerIndex;
		this.player = null;

		this.rootElem.innerHTML = `
			<div class="player-label">
				<span class="player-name"></span>
			</div>
			<div class="player-options">
			</div>
		`;

		this.nameElem = this.rootElem.getElementsByClassName('player-name')[0] as HTMLSpanElement;

		this.party.changeEmitter.on(() => {
			const newPlayer = this.party.getPlayer(this.playerIndex);

			if (((newPlayer == null) != (this.player == null)) || newPlayer != this.player) {
				this.setPlayer(newPlayer);
				return;
			}

			this.update();
		});
	}

	setPlayer(newPlayer: Player<any> | null) {
		this.player = newPlayer;
		this.party.setPlayer(this.playerIndex, this.player);

		this.update();
	}

	private update() {
		if (this.player == null) {
			this.nameElem.textContent = '';
		} else {
			this.nameElem.textContent = this.player.getName();
		}
	}
}

export class NewPlayerPicker extends Component {
	private readonly raid: Raid;
	private currentFaction: Faction;

  constructor(parent: HTMLElement, raid: Raid, specs: Array<PresetSpecSettings<any>>) {
    super(parent, 'new-player-picker-root');
		this.raid = raid;
		this.currentFaction = Faction.Alliance;

		this.rootElem.innerHTML = `
			<div class="faction-selector"></div>
			<div class="phase-selector"></div>
			<div class="class-pickers"></div>
		`;

		const factionSelector = new EnumPicker<NewPlayerPicker>(this.rootElem.getElementsByClassName('faction-selector')[0] as HTMLElement, this, {
			label: 'Faction',
			labelTooltip: 'Default faction for newly-created players.',
			values: [
				{ name: 'Alliance', value: Faction.Alliance },
				{ name: 'Horde', value: Faction.Horde },
			],
			changedEvent: (picker: NewPlayerPicker) => new TypedEvent<void>(),
			getValue: (picker: NewPlayerPicker) => picker.currentFaction,
			setValue: (picker: NewPlayerPicker, newValue: number) => {
				picker.currentFaction = newValue;
			},
		});

		const phaseSelector = makePhaseSelector(this.rootElem.getElementsByClassName('phase-selector')[0] as HTMLElement, this.raid.sim);
	}
}

export interface PresetSpecSettings<SpecType extends Spec> {
	spec: Spec,
	rotation: SpecRotation<SpecType>,
	talents: string,
	specOptions: SpecOptions<SpecType>,
	consumes: Consumes,

	defaultFactionRaces: Record<Faction, Race>,
	defaultGear: Record<Faction, Record<number, EquipmentSpec>>,

	tooltip: string,
	iconUrl: string,
}
