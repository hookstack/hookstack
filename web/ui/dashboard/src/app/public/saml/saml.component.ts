import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoaderModule } from 'src/app/private/components/loader/loader.module';
import { SamlService } from './saml.service';
import { ActivatedRoute, Router } from '@angular/router';
import { PrivateService } from 'src/app/private/private.service';

@Component({
	selector: 'convoy-saml',
	standalone: true,
	imports: [CommonModule, LoaderModule],
	templateUrl: './saml.component.html',
	styleUrls: ['./saml.component.scss']
})
export class SamlComponent implements OnInit {
	accessCode: string = this.route.snapshot.queryParams.saml_access_code;

	constructor(private samlService: SamlService, private route: ActivatedRoute, private router: Router, private privateService: PrivateService) {}

	ngOnInit() {
		this.authenticate();
	}

	async authenticate() {
		try {
			const response = await this.samlService.authenticateWithSaml(this.accessCode);

			localStorage.setItem('CONVOY_AUTH', JSON.stringify(response.data));
			localStorage.setItem('CONVOY_AUTH_TOKENS', JSON.stringify(response.data.token));

			await this.getOrganisations();
			this.router.navigateByUrl('/');
		} catch {
            this.router.navigateByUrl('/login');
        }
	}

    async getOrganisations() {
		try {
			await this.privateService.getOrganizations({ refresh: true });
			return;
		} catch (error) {
			return error;
		}
	}
}
