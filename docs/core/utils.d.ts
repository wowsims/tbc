export declare function equalsOrBothNull<T>(a: T, b: T, comparator?: (_a: NonNullable<T>, _b: NonNullable<T>) => boolean): boolean;
export declare function sum(arr: Array<number>): number;
export declare function intersection<T>(a: Array<T>, b: Array<T>): Array<T>;
export declare function wait(ms: number): Promise<void>;
export declare function getEnumValues<E>(enumType: any): Array<E>;
export declare function isRightClick(event: MouseEvent): boolean;
