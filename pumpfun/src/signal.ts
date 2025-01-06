export class Signal {
  private _done = false;

  constructor(...events: NodeJS.Signals[]) {
    for (const event of events) {
      process.on(event, () => {
        this._done = true;
      });
    }
  }

  get done() {
    return this._done;
  }
}
