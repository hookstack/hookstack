import { Component, HostListener, OnInit, ViewChild } from '@angular/core';
import { PROJECT } from 'src/app/models/project.model';
import { PrivateService } from '../../private.service';
import { Router } from '@angular/router';
import { LicensesService } from 'src/app/services/licenses/licenses.service';

@Component({
	selector: 'app-project',
	templateUrl: './project.component.html',
	styleUrls: ['./project.component.scss']
})
export class ProjectComponent implements OnInit {
	screenWidth = window.innerWidth;

	sideBarItems = [
		{
			name: 'Event Deliveries',
			icon: 'events',
			route: '/events'
		},
		{
			name: 'Sources',
			icon: 'sources',
			route: '/sources'
		},
		{
			name: 'Subscriptions',
			icon: 'subscriptions',
			route: '/subscriptions'
		},
		{
			name: 'Endpoints',
			icon: 'endpoint',
			route: '/endpoints'
		},
		{
			name: 'Events Log',
			icon: 'logs',
			route: '/events-log'
		},
		{
			name: 'Meta Events',
			icon: 'meta',
			route: '/meta-events'
		}
	];
	secondarySideBarItems = [
		{
			name: 'Events Log',
			icon: 'logs',
			route: '/events-log'
		},
		{
			name: 'Meta Events',
			icon: 'meta',
			route: '/meta-events'
		}
	];
	shouldShowFullSideBar = true;
	projectDetails?: PROJECT;
	isLoadingProjectDetails: boolean = true;
	showHelpDropdown = false;
	projects: PROJECT[] = [];
	activeNavTab: any;

	constructor(private privateService: PrivateService, private router: Router, public licenseService: LicensesService) {}

	ngOnInit() {
		Promise.all([this.getProjectDetails(), this.getProjects()]);
	}

	get activeTab(): any {
		const element = document.querySelector('.nav-tab.on') as any;
		if (element) this.activeNavTab = element;
		return element || this.activeNavTab;
	}

	async getProjectDetails() {
		this.isLoadingProjectDetails = true;

		try {
			const projectDetails = await this.privateService.getProjectDetails;
			this.projectDetails = projectDetails;
			if (this.projectDetails?.type === 'outgoing') this.sideBarItems.push({ name: 'Portal Links', icon: 'portal', route: '/portal-links' });
			this.isLoadingProjectDetails = false;
		} catch (error) {
			this.isLoadingProjectDetails = false;
		}
	}

	async getProjects() {
		try {
			const response = await this.privateService.getProjects();
			this.projects = response.data;
		} catch (error) {}
	}

	isOutgoingProject(): boolean {
		return this.projectDetails?.type === 'outgoing';
	}

	isStrokeIcon(icon: string): boolean {
		const menuIcons = ['subscriptions', 'portal', 'logs', 'meta'];
		const checkForStrokeIcon = menuIcons.some(menuIcon => icon.includes(menuIcon));

		return checkForStrokeIcon;
	}

	getFirstletters(text?: string) {
		if (!text) return;
		const firstLetters = text
			.split(' ')
			.map(word => word[0])
			.join('')
			.slice(0, 2);

		return firstLetters;
	}

	async getProjectCompleteDetails(project: PROJECT) {
		this.isLoadingProjectDetails = true;

		try {
			this.projectDetails = project;
			localStorage.setItem('CONVOY_PROJECT', JSON.stringify(this.projectDetails));

			if (this.projectDetails?.type === 'outgoing' && this.sideBarItems[this.sideBarItems.length - 1].icon === 'endpoint') this.sideBarItems.push({ name: 'Portal Links', icon: 'portal', route: '/portal-links' });
			if (this.projectDetails?.type === 'incoming' && this.sideBarItems[this.sideBarItems.length - 1].icon === 'portal') this.sideBarItems.pop();

			await this.privateService.getProject({ refresh: true, projectId: project.uid });
			await this.privateService.getProjectStat({ refresh: true });
			this.router.navigateByUrl('/', { skipLocationChange: true }).then(() => {
				this.router.navigate([`/projects/${project.uid}`]);
			});

			this.isLoadingProjectDetails = false;
		} catch (error) {
			this.isLoadingProjectDetails = false;
		}
	}

	checkScreenSize() {
		this.screenWidth > 1150 ? (this.shouldShowFullSideBar = true) : (this.shouldShowFullSideBar = false);
	}

	@HostListener('window:resize', ['$event'])
	onWindowResize() {
		this.screenWidth = window.innerWidth;
		this.checkScreenSize();
	}
}
