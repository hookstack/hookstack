import { Location } from '@angular/common';
import { Component, HostListener, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { APP, ENDPOINT } from 'src/app/models/app.model';
import { PAGINATION } from 'src/app/models/global.model';
import { PrivateService } from 'src/app/private/private.service';
import { GeneralService } from 'src/app/services/general/general.service';
import { AppDetailsService } from './app-details.service';
import { CliKeysComponent } from './cli-keys/cli-keys.component';

@Component({
	selector: 'app-app-details',
	templateUrl: './app-details.component.html',
	styleUrls: ['./app-details.component.scss']
})
export class AppDetailsComponent implements OnInit {
	@ViewChild(CliKeysComponent) cliKeys!: CliKeysComponent;
	showAddEndpointModal = false;
	showAddEventModal = false;
	showEndpointSecret = false;
	isSendingNewEvent = false;
	savingEndpoint = false;
	loadingAppPotalToken = false;
	isLoadingAppDetails = false;
	shouldRenderSmallSize = false;
	showDeleteModal = false;
	isDeletingEndpoint = false;
	screenWidth = window.innerWidth;
	appPortalLink!: string;
	appPortalIframe!: string;
	endpointSecretKey!: string;
	appId!: string;
	appsDetailsItem!: APP;
	apps!: { pagination: PAGINATION; content: APP[] };
	selectedEndpoint?: ENDPOINT;
	tabs: ['cli keys', 'devices'] = ['cli keys', 'devices'];
	activeTab: 'cli keys' | 'devices' = 'cli keys';
	isCliAvailable: boolean = false;

	constructor(private appDetailsService: AppDetailsService, private generalService: GeneralService, private route: ActivatedRoute, private location: Location, private router: Router, public privateService: PrivateService) {}

	async ngOnInit() {
		this.isLoadingAppDetails = true;
		this.isCliAvailable = await this.privateService.getFlag('can_create_cli_api_key');
		if (this.privateService.activeProjectDetails?.type === 'outgoing') this.loadingAppPotalToken = true;
		this.checkScreenSize();
		this.getAppDetails(this.route.snapshot.params.id);
	}

	goBack() {
		this.location.back();
	}

	viewEndpointSecretKey(secretKey: string) {
		this.showEndpointSecret = !this.showEndpointSecret;
		this.endpointSecretKey = secretKey;
	}

	async getAppDetails(appId: string) {
		this.selectedEndpoint = undefined;
		this.isLoadingAppDetails = true;

		try {
			const response = await this.appDetailsService.getApp(appId);
			this.appsDetailsItem = response.data;
			this.getAppPortalToken({ redirect: false });
			this.isLoadingAppDetails = false;
		} catch {
			this.isLoadingAppDetails = false;
		}
	}

	async getAppPortalToken(requestDetail: { redirect: boolean }) {
		if (this.privateService.activeProjectDetails?.type === 'incoming') return;

		this.loadingAppPotalToken = true;
		const payload = {
			key_type: 'app_portal'
		};
		try {
			const appTokenResponse = await this.appDetailsService.generateKey({ appId: this.appsDetailsItem.uid, body: payload });
			this.appPortalLink = appTokenResponse.data.url;
			this.appPortalIframe = `<iframe style="width: 100%; height: 98%; border: none;" frameborder="0" src="${appTokenResponse.data.url}"></iframe>`;
			if (requestDetail.redirect) window.open(`${appTokenResponse.data.url}`, '_blank');
			this.loadingAppPotalToken = false;
		} catch (error) {
			this.loadingAppPotalToken = false;
			return error;
		}
	}

	loadEventsFromAppsTable(appId: string) {
		const projectId = this.appDetailsService.projectId;
		this.router.navigate(['/projects/' + projectId + '/events'], { queryParams: { eventsApp: appId } });
	}

	async deleteEndpoint() {
		this.isDeletingEndpoint = true;
		try {
			const response = await this.appDetailsService.deleteEndpoint({ appId: this.appsDetailsItem?.uid, endpointId: this.selectedEndpoint?.uid || '' });
			this.generalService.showNotification({ style: 'success', message: response.message });
			this.showDeleteModal = false;
			this.isDeletingEndpoint = false;
			this.getAppDetails(this.appsDetailsItem?.uid);
		} catch {
			this.isDeletingEndpoint = false;
		}
	}

	toggleActiveTab(tab: 'cli keys' | 'devices') {
		this.activeTab = tab;
	}

	checkScreenSize() {
		this.screenWidth > 1010 ? (this.shouldRenderSmallSize = false) : (this.shouldRenderSmallSize = true);
	}

	closeEditEndpointModal() {
		this.showAddEndpointModal = false;
		this.selectedEndpoint = undefined;
	}

	focusInput() {
		document.getElementById('tagInput')?.focus();
	}

	@HostListener('window:resize', ['$event'])
	onWindowResize() {
		this.screenWidth = window.innerWidth;
		this.checkScreenSize();
	}
}
