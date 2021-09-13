// Returns a new array containing only elements present in both a and b.
export function intersection<T>(a: Array<T>, b: Array<T>): Array<T> {
  return a.filter(value => b.includes(value));
}

export async function wait(ms: number): Promise<void> {
  return new Promise(resolve => setTimeout(resolve, ms));
}

// Only works for numeric enums
export function getEnumValues<E>(enumType: any): Array<E> {
  return Object.keys(enumType)
      .filter(key => !isNaN(Number(enumType[key])))
      .map(key => parseInt(enumType[key]) as unknown as E);
}

// Whether a click event was a right click.
export function isRightClick(event: MouseEvent): boolean {
  return event.button == 2;
}
