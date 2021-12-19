


export interface BuffBotSettings {
	// The value of this field must never change, to preserve local storage data.
	buffBotId: string,

	spec: Spec,
	name: string,
	tooltip: string,
	iconUrl: string,

	// Callback to apply buffs from this buff bot.
	modifyRaidProto: (raidProto: RaidProto, partyProto: PartyProto) => void,
	modifyEncounterProto: (encounterProto: EncounterProto) => void,
}

// Represents a buff bot in a raid.
export interface BuffBotData {
	readonly settings: BuffBotSettings;

	private raidIndex: number;

	private innervateAssignment: number;
	private powerInfusionAssignment: number;
};

export class BuffBot {

}
