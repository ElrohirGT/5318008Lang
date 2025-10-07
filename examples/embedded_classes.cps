class Legs {
	let count: integer;

	function constructor(legCount: integer) {
		this.count = legCount;
	}
}

class Dog {
	let isAlive: boolean = true;
	let legs: Legs = new Legs(4);
	let age: integer;

	function constructor(age: integer, isAlive: boolean) {
		this.age = age;
		this.isAlive = isAlive;
	}
}

let a = new Dog(8, false);