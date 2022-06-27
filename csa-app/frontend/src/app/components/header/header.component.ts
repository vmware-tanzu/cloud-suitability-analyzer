import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import '@cds/core/icon/register.js';
import {AnalyzerRunService} from '../../services/analyzerrun.service';
import {pushErrorNotification} from '../../utils/notificationutil';
import {ToastrService} from 'ngx-toastr';
import { ClarityIcons } from '@cds/core/icon';

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
    const csaSvg = `<svg width="36" height="36" viewBox="0 0 36 36" fill="none" xmlns="http://www.w3.org/2000/svg">
<path d="M32.0308 23.6013C32.0308 23.7313 32.0308 23.8613 32.0308 23.9913C32.0983 26.0625 30.4815 27.8 28.4108 27.8813H8.95077C6.02388 27.7781 3.73087 25.3287 3.82077 22.4013C3.74767 19.4856 6.0359 17.0546 8.95077 16.9513H10.5808L10.6708 16.0713C11.1305 11.6199 14.8175 8.19816 19.2908 8.07133H19.3708C21.943 8.11746 24.3555 9.32741 25.9308 11.3613H28.3008C26.451 8.17583 23.0736 6.18582 19.3908 6.11133H19.2808C14.1621 6.2286 9.82635 9.91754 8.89077 14.9513C4.92102 15.1032 1.80601 18.4096 1.89077 22.3813C1.80128 26.3854 4.96705 29.7077 8.97077 29.8113H28.4408C29.958 29.7747 31.3984 29.1365 32.4449 28.0375C33.4915 26.9384 34.0584 25.4685 34.0208 23.9513C34.0275 23.8214 34.0275 23.6912 34.0208 23.5613L32.0308 23.6013Z" fill="#0091DA"/>
<path d="M28.6659 14.6115C28.6786 14.2476 28.5445 13.8941 28.2944 13.6322C28.0443 13.3703 27.6998 13.2224 27.3398 13.2227H27.2618C26.5036 13.2227 25.8889 13.8445 25.8889 14.6115C25.8889 15.3786 26.5036 16.0004 27.2618 16.0004C28.0219 15.9924 28.6408 15.3801 28.6659 14.6115Z" fill="#EDB200"/>
<rect x="19.8889" y="23.0625" width="2.44444" height="2.44443" fill="#F54F47"/>
<path d="M8.55557 22.8403L10.1914 25.3403H6.91975L8.55557 22.8403Z" fill="#EDB200"/>
<path d="M15 13L16.3031 13.6275L16.6249 15.0375L15.7232 16.1683H14.2769L13.3752 15.0375L13.697 13.6275L15 13Z" fill="#60B515"/>
<path fill-rule="evenodd" clip-rule="evenodd" d="M34.0174 16.3022L27.4445 22.8752L21.7778 17.2085L14.5 24.4863L10.5937 20.58L11.8508 19.3229L14.5 21.9721L21.7778 14.6943L27.4445 20.361L32.7604 15.0451L34.0174 16.3022Z" fill="#00C1D5"/>
</svg>`;
    ClarityIcons.addIcons(['csa', csaSvg]);
  }

  ngOnInit(): void {
    this.analyzerRunService.getForgeVersion().subscribe(version => {
      /* istanbul ignore else */
      if (version) {
        Object.keys(version).forEach((key) => {
          this.forgeVersion = version[key];
        });
      }
    }, error => {
      pushErrorNotification(error, this.toastr);
    });
  }

}
