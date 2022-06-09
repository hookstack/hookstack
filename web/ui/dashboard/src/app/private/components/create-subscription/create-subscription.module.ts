import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateSubscriptionComponent } from './create-subscription.component';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAppModule } from '../create-app/create-app.module';
import { CreateSourceModule } from '../create-source/create-source.module';
import { TooltipModule } from '../tooltip/tooltip.module';

@NgModule({
	declarations: [CreateSubscriptionComponent],
	imports: [CommonModule, ReactiveFormsModule, CreateAppModule, CreateSourceModule, TooltipModule],
	exports: [CreateSubscriptionComponent]
})
export class CreateSubscriptionModule {}
