import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import { cloudIcon } from '@cds/core/icon';
import '@cds/core/icon/register.js';
import {AnalyzerRunService} from '../../services/analyzerrun.service';
import { pushErrorNotification } from '../../utils/notificationutil';
import {ToastrService} from "ngx-toastr";
import { ClarityIcons } from '@clr/icons';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  @Input()
  selectedRunId: number | undefined;

  @Output()
  selectedRunIdChange = new EventEmitter();

  forgeVersion: string = 'v3.2.5-rc1';

  constructor(private analyzerRunService: AnalyzerRunService, public toastr: ToastrService) {
    const csaSvg = `<svg width="324" height="324" viewBox="0 0 324 324" fill="none" xmlns="http://www.w3.org/2000/svg">
    <path d="M288.277 212.41C288.277 213.58 288.277 214.75 288.277 215.92C288.885 234.561 274.333 250.198 255.697 250.93H80.5567C54.2147 250.001 33.5776 227.956 34.3867 201.61C33.7288 175.368 54.3229 153.489 80.5567 152.56H95.2267L96.0367 144.64C100.174 104.577 133.357 73.7815 173.617 72.64H174.337C197.486 73.0552 219.199 83.9447 233.377 102.25H254.707C238.059 73.5805 207.663 55.6704 174.517 55H173.527C127.459 56.0554 88.4369 89.2559 80.0167 134.56C44.2889 135.927 16.2538 165.684 17.0167 201.43C16.2113 237.467 44.7032 267.367 80.7367 268.3H255.967C269.621 267.97 282.585 262.227 292.004 252.335C301.423 242.444 306.525 229.215 306.187 215.56C306.247 214.391 306.247 213.219 306.187 212.05L288.277 212.41Z" fill="#0091DA"/>
    <path d="M257.993 131.5C258.107 128.225 256.9 125.043 254.649 122.686C252.399 120.328 249.298 118.998 246.058 119H245.356C238.532 119 233 124.596 233 131.5C233 138.404 238.532 144 245.356 144C252.197 143.928 257.767 138.417 257.993 131.5Z" fill="#EDB200"/>
    <rect x="179" y="207.561" width="22" height="22" fill="#F54F47"/>
    <path d="M77 205.561L91.7224 228.061H62.2776L77 205.561Z" fill="#EDB200"/>
    <path d="M135 117L146.727 122.648L149.624 135.338L141.508 145.515H128.492L120.376 135.338L123.273 122.648L135 117Z" fill="#60B515"/>
    <path fill-rule="evenodd" clip-rule="evenodd" d="M306.156 146.718L247 205.874L196 154.874L130.5 220.374L95.3428 185.218L106.656 173.904L130.5 197.747L196 132.247L247 183.247L294.843 135.404L306.156 146.718Z" fill="#00C1D5"/>
    </svg>`;
    ClarityIcons.add({'csa': csaSvg});
  }

  ngOnInit(): void {
    this.analyzerRunService.getForgeVersion().subscribe(version => {
      if (version && version.length > 0) {
        this.forgeVersion = version;
      }
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

}
