import { StatusColorPipe } from '../status-color/status-color.pipe';

describe('StatusColorPipe', () => {
	it('create an instance', () => {
		const pipe = new StatusColorPipe();
		expect(pipe).toBeTruthy();
	});
});
