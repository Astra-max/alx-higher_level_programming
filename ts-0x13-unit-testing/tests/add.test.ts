import {add} from "../src/add";


describe("Testing utility functions", ()=> {
	test("add two numbers func", () => {
		expect(add(2,4)).toBe(6)
	});
});
