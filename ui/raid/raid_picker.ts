import { Component } from '/tbc/core/components/component.js';
import { Raid } from '/tbc/core/raid.js';
import { MAX_PARTY_SIZE } from '/tbc/core/party.js';
import { Party } from '/tbc/core/party.js';
import { Player } from '/tbc/core/player.js';

export class RaidPicker extends Component {
	private readonly raid: Raid;
  private readonly partyPickers: Array<PartyPicker>;

  constructor(parent: HTMLElement, raid: Raid) {
    super(parent, 'raid-picker-root');
		this.raid = raid;

    const raidViewer = document.createElement('div');
    raidViewer.classList.add('current-raid-viewer');
    this.rootElem.appendChild(raidViewer);
		raidViewer.innerHTML = `
			<div class="parties-container">
			</div>
		`;

    const newPlayerPicker = document.createElement('div');
    newPlayerPicker.classList.add('new-player-picker');
    this.rootElem.appendChild(newPlayerPicker);

    const partiesContainer = this.rootElem.getElementsByClassName('parties-container')[0] as HTMLDivElement;
		this.partyPickers = this.raid.getParties().map((party, i) => new PartyPicker(partiesContainer, party, i));
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
		this.playerPickers = [...Array(MAX_PARTY_SIZE).keys()].map(i => new PlayerPicker(playersContainer, i));
	}
}

export class PlayerPicker extends Component {
	private readonly party: Party;
	private readonly playerIndex: number;
	private player: Player?;

  constructor(parent: HTMLElement, party: Party, playerIndex: number) {
    super(parent, 'player-picker-root');
		this.party = party;
		this.playerIndex = playerIndex;

		this.rootElem.innerHTML = `
			<div class="player-label">
				<span></span>
			</div>
			<div class="player-options">
			</div>
		`;
	}

	setPlayer(newPlayer: Player?) {
		this.player = newPlayer;
	}
}
