import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { PAGINATION } from 'src/app/models/global.model';
import { SUBSCRIPTION } from 'src/app/models/subscription';
import { PrivateService } from 'src/app/private/private.service';
import { SubscriptionsService } from './subscriptions.service';

@Component({
	selector: 'app-subscriptions',
	templateUrl: './subscriptions.component.html',
	styleUrls: ['./subscriptions.component.scss']
})
export class SubscriptionsComponent implements OnInit {
	activeSubscription?: SUBSCRIPTION;
	shouldShowCreateSubscriptionModal = this.router.url.split('/')[4] === 'new';
	projectId!: string;
	subscriptions!: { content: SUBSCRIPTION[]; pagination: PAGINATION };

	constructor(private route: ActivatedRoute, private privateService: PrivateService, private router: Router, private subscriptionsService: SubscriptionsService) {
		this.projectId = this.privateService.activeProjectId;
	}

	async ngOnInit() {
		await this.getSubscriptions();

		this.route.queryParams.subscribe(params => (this.activeSubscription = this.subscriptions?.content.find(source => source.uid === params?.id)));
	}

	async getSubscriptions(requestDetails?: { page?: number }) {
		try {
			const subscriptionsResponse = await this.subscriptionsService.getSubscriptions({ page: requestDetails?.page });
			this.subscriptions = subscriptionsResponse.data;
		} catch (error) {
			console.log(error);
		}
	}
}
