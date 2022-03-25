import { Consumes } from '/tbc/core/proto/common.js';
import { Flask } from '/tbc/core/proto/common.js';
import { Food } from '/tbc/core/proto/common.js';
import { Potions } from '/tbc/core/proto/common.js';
import { ProtectionWarrior_Rotation as ProtectionWarriorRotation, ProtectionWarrior_Options as ProtectionWarriorOptions } from '/tbc/core/proto/warrior.js';
// Preset options for this spec.
// Eventually we will import these values for the raid sim too, so its good to
// keep them in a separate file.
// Default talents. Uses the wowhead calculator format, make the talents on
// https://tbc.wowhead.com/talent-calc and copy the numbers in the url.
export const ImpaleProtTalents = {
    name: 'Impale Prot',
    data: '35000301302-03-0055511033010101501351',
};
export const DefaultRotation = ProtectionWarriorRotation.create({});
export const DefaultOptions = ProtectionWarriorOptions.create({
    startingRage: 0,
    precastSapphire: false,
    precastT2: false,
});
export const DefaultConsumes = Consumes.create({
    flask: Flask.FlaskOfRelentlessAssault,
    food: Food.FoodRoastedClefthoof,
    defaultPotion: Potions.HastePotion,
});
