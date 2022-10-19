import {browser, by, element, ExpectedConditions} from 'protractor';
import {ElementArrayFinder} from "protractor/built/element";

const fs = require('fs');
const path = require('path');

const downloadsPath = path.resolve(__dirname, '../../../build/downloads');

export class DownloadHelper {
  private readonly maxReadWaitTime = 20000; // 10 seconds

  constructor() {
  }

  removeFile(fileName: string): void {
    const filePath = this.getFilePath(fileName);
    if (fs.existsSync(filePath)) {
      fs.unlinkSync(filePath);
    }
  }

  async getFileContents(fileName: string): Promise<string> {
    const filePath = this.getFilePath(fileName);
    try {
      await browser.driver.wait(() => fs.existsSync(filePath), this.maxReadWaitTime);
    } catch (err) {
      throw new Error(`File ${filePath} does not exist`);
    }

    return fs.readFileSync(filePath, { encoding: 'utf8' });
  }

  private getFilePath(fileName: string): string {
    return path.join(downloadsPath, fileName);
  }

  getDownloadFileNameForExecutiveSummary(): string {
    return 'app-summary.xlsx';
  }
}
