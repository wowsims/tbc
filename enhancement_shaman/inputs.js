import { EnhancementShaman_Rotation_RotationType as RotationType } from '/tbc/core/proto/shaman.js';
// Configuration for spec-specific UI elements on the settings tab.
// These don't need to be in a separate file but it keeps things cleaner.
export const IconBloodlust = makeBooleanShamanBuffInput({ spellId: 2825 }, 'bloodlust');
// export const IconManaSpringTotem = makeBoolShamanTotem({ spellId: 25570 }, 'manaSpringTotem');
// export const IconTotemOfWrath = makeBoolShamanTotem({ spellId: 30706 }, 'totemOfWrath');
export const IconWaterShield = makeBooleanShamanBuffInput({ spellId: 33736 }, 'waterShield');
// export const IconWrathOfAirTotem = makeBoolShamanTotem({ spellId: 3738 }, 'wrathOfAirTotem');
export const EnhancementShamanRotationConfig = {
    inputs: [
        {
            type: 'enum', cssClass: 'rotation-enum-picker',
            getModObject: (simUI) => simUI.player,
            config: {
                label: 'Type',
                values: [
                    {
                        name: 'Basic', value: RotationType.Basic,
                        tooltip: 'does basic stuff',
                    },
                ],
                changedEvent: (player) => player.rotationChangeEmitter,
                getValue: (player) => player.getRotation().type,
                setValue: (player, newValue) => {
                    const newRotation = player.getRotation();
                    newRotation.type = newValue;
                    player.setRotation(newRotation);
                },
            },
        }
    ],
};
function makeBooleanShamanBuffInput(id, optionsFieldName) {
    return {
        id: id,
        states: 2,
        changedEvent: (player) => player.specOptionsChangeEmitter,
        getValue: (player) => player.getSpecOptions()[optionsFieldName],
        setValue: (player, newValue) => {
            const newOptions = player.getSpecOptions();
            newOptions[optionsFieldName] = newValue;
            player.setSpecOptions(newOptions);
        },
    };
}
// function makeBoolShamanTotem(id: ItemOrSpellId, optionsFieldName: keyof totems?): IconPickerConfig<Player<any>, boolean> {
//   return {
//     id: id,
//     states: 2,
// 		changedEvent: (player: Player<Spec.SpecEnhancementShaman>) => player.specOptionsChangeEmitter,
// 		getValue: (player: Player<Spec.SpecEnhancementShaman>) => {
// 			const totems = player.getSpecOptions().totems as ShamanTotems;
// 			return totems[optionsFieldName] as boolean;
// 		},
// 		setValue: (player: Player<Spec.SpecEnhancementShaman>, newValue: boolean) => {
// 			const newOptions = player.getSpecOptions();
// 			const totems = newOptions.totems as ShamanTotems;
//       		(totems[optionsFieldName] as boolean) = newValue;
// 			newOptions.totems = totems;
// 			player.setSpecOptions(newOptions);
// 		},
//   }
// }
