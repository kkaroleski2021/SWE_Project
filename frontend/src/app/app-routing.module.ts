import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ListingPageComponent } from './listing-page/listing-page.component';
import { ProfilesComponent } from './profiles/profiles.component';
import { SellerPageComponent } from './seller-page/seller-page.component';
import { SComponent } from './signup/signup.component';
import { BComponent} from './login/login.component';
import { OrderFormComponent } from './order-form/order-form.component';
import { HomeComponent } from './home/home.component';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent
  },
  {
    path: 'Login',
    component: BComponent
  },
  {
    path: 'Signup',
    component: SComponent
  },
  {
    path: 'ListingPage',
    component: ListingPageComponent
  },
  {
    path: 'SellerPage',
    component: SellerPageComponent
  },
  {
    path: 'OrderForm',
    component: OrderFormComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
  static withRoutes(arg0: { path: string; component: typeof HomeComponent; }[]): any {
    throw new Error('Method not implemented.');
  }
}
