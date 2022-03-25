import { Component } from '/tbc/core/components/component.js';
import { IconEnumPicker } from '/tbc/core/components/icon_enum_picker.js';
import { EventID, TypedEvent } from '/tbc/core/typed_event.js';
import { Class } from '/tbc/core/proto/common.js';
import { Spec } from '/tbc/core/proto/common.js';
import { Blessings } from '/tbc/core/proto/ui.js';
import { BlessingsAssignment } from '/tbc/core/proto/ui.js';
import { BlessingsAssignments } from '/tbc/core/proto/ui.js';
import { ActionId } from '/tbc/core/proto_utils/action_id.js';
import {
	makeDefaultBlessings,
	classColors,
	naturalSpecOrder,
	specNames,
	titleIcons,
} from '/tbc/core/proto_utils/utils.js';
import { getEnumValues } from '/tbc/core/utils.js';

import { RaidSimUI } from './raid_sim_ui.js';
import { implementedSpecs } from './presets.js';

declare var tippy: any;

const MAX_PALADINS = 4;

export class BlessingsPicker extends Component {
	readonly raidSimUI: RaidSimUI;
	readonly changeEmitter: TypedEvent<void> = new TypedEvent<void>();

	private readonly rows: Array<HTMLTableRowElement>;

	private assignments: BlessingsAssignments;

	constructor(parentElem: HTMLElement, raidSimUI: RaidSimUI) {
		super(parentElem, 'blessings-picker-root');
		this.raidSimUI = raidSimUI;
		this.assignments = BlessingsAssignments.clone(makeDefaultBlessings(4));

		this.rootElem.innerHTML = `
		<table class="blessings-table">
			<thead class="blessings-table-header">
				<tr class="blessings-table-header-row">
					<th class="blessings-table-header-cell"></th>
				</tr>
			</thead>
			<tbody class="blessings-table-body">
			</tbody>
		</table>
		`;

		const headerRow = this.rootElem.getElementsByClassName('blessings-table-header-row')[0] as HTMLTableRowElement;
		const bodyElem = this.rootElem.getElementsByClassName('blessings-table-body')[0] as HTMLTableSectionElement;

		naturalSpecOrder.forEach(spec => {
			if (!implementedSpecs.includes(spec)) {
				return;
			}

			const cell = document.createElement('th');
			cell.classList.add('blessings-table-header-cell');
			headerRow.appendChild(cell);

			const icon = document.createElement('img');
			icon.src = titleIcons[spec];
			cell.appendChild(icon);

			tippy(icon, {
				'content': specNames[spec],
				'allowHTML': true,
			});
		});

		this.rows = [...Array(MAX_PALADINS).keys()].map(rowIndex => {
			const row = document.createElement('tr');
			row.classList.add('blessings-table-row');
			bodyElem.appendChild(row);

			const cell = document.createElement('td');
			cell.classList.add('blessings-table-cell', 'blessings-table-label-cell');
			cell.textContent = 'Paladin ' + (rowIndex + 1);
			row.appendChild(cell);

			naturalSpecOrder.forEach(spec => {
				if (!implementedSpecs.includes(spec)) {
					return;
				}

				const cell = document.createElement('td');
				cell.classList.add('blessings-table-cell');
				row.appendChild(cell);
				const blessingPicker = new IconEnumPicker(cell, this, {
					extraCssClasses: [
						'blessing-picker',
					],
					numColumns: 1,
					values: [
						{ color: classColors[Class.ClassPaladin], value: Blessings.BlessingUnknown },
						{ actionId: ActionId.fromSpellId(25898), value: Blessings.BlessingOfKings },
						{ actionId: ActionId.fromSpellId(25895), value: Blessings.BlessingOfSalvation },
						{ actionId: ActionId.fromSpellId(27141), value: Blessings.BlessingOfMight },
						{ actionId: ActionId.fromSpellId(27143), value: Blessings.BlessingOfWisdom },
					],
					equals: (a: Blessings, b: Blessings) => a == b,
					zeroValue: Blessings.BlessingUnknown,
					changedEvent: (picker: BlessingsPicker) => picker.changeEmitter,
					getValue: (picker: BlessingsPicker) => picker.assignments.paladins[rowIndex]?.blessings[spec] || Blessings.BlessingUnknown,
					setValue: (eventID: EventID, picker: BlessingsPicker, newValue: number) => {
						const currentValue = picker.assignments.paladins[rowIndex].blessings[spec];
						if (currentValue != newValue) {
							picker.assignments.paladins[rowIndex].blessings[spec] = newValue;
							this.changeEmitter.emit(eventID);
						}
					},
				});
			});

			return row;
		});

		this.setNumPaladins(raidSimUI.getClassCount(Class.ClassPaladin));
		raidSimUI.compChangeEmitter.on(eventID => {
			this.setNumPaladins(raidSimUI.getClassCount(Class.ClassPaladin));
		});
	}

	private setNumPaladins(numPaladins: number) {
		numPaladins = Math.min(numPaladins, MAX_PALADINS);

		for (let i = 0; i < numPaladins; i++) {
			this.rows[i].classList.add('paladin-active');
		}
		for (let i = numPaladins; i < this.rows.length; i++) {
			this.rows[i].classList.remove('paladin-active');
		}
	}

	getAssignments(): BlessingsAssignments {
		// Defensive copy.
		return BlessingsAssignments.clone(this.assignments);
	}

	setAssignments(eventID: EventID, newAssignments: BlessingsAssignments) {
		this.assignments = BlessingsAssignments.clone(newAssignments);
		this.changeEmitter.emit(eventID);
	}
}
