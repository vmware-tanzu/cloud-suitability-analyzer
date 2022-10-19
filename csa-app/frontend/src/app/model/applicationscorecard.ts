
export class ApplicationScoreCard {
    application: string;
    info: number;
    low: number;
    medium: number;
    high: number;
    total: number;
  
    constructor(application: string, info: number, low: number, medium: number, high: number, total: number) {
      this.application = application;
      this.info = info;
      this.low = low;
      this.medium = medium;
      this.high = high;
      this.total = total;
    }
  }
  
  
  