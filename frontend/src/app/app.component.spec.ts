import { TestBed } from '@angular/core/testing';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { FooterComponent } from './footer/footer.component';
import { HeaderComponent } from './header/header.component';
import { HomeComponent } from './home/home.component';
import { ListingPageComponent } from './listing-page/listing-page.component';
import { BComponent} from './login/login.component';
import { SComponent } from './signup/signup.component';
import { Header2Component } from './loginHeader/loginHeader.component'; 


describe('AppComponent', () => {
beforeEach(async () => {
await TestBed.configureTestingModule({
declarations: [
AppComponent
],
}).compileComponents();
});

it(`should have as title 'SimpleApp'`, () => {
const fixture = TestBed.createComponent(AppComponent);
const app = fixture.componentInstance;
expect(app.title).equal('SimpleApp');
});

it('should render title', () => {
const fixture = TestBed.createComponent(AppComponent);
fixture.detectChanges();
const compiled = fixture.nativeElement as HTMLElement;
expect(compiled.querySelector('.content span')?.textContent).contain('SimpleApp app is running!');

});

beforeEach(() => {
  TestBed.configureTestingModule({
    imports: [
        AppRoutingModule.withRoutes(
        [{path: '', component: HomeComponent}, {path: 'Login', component: BComponent}, {path: 'signup', component: SComponent}]
      )
    ]
  });

})

it('should render title', () => { //6
    const fixture = TestBed.createComponent(Header2Component);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector('h1').textContent).contain('Log In to Swampy Sells!');
  });

  it(`should have as title 'hello-world'`, () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app.title).equal('hello-world');
  });

  it('should set image path as expected', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const ele = fixture.debugElement.nativeElement.querySelectorAll('img');
    expect(ele[1]['src']).contain('assets/images/stuffedAnimalChair.jpg');
    expect(ele[2]['src']).contain('assets/images/clamshellVase.jpg');
    expect(ele[3]['src']).contain('assets/images/grandmasCouch.jpg');
    expect(ele[4]['src']).contain('assets/images/clamshellVase.jpg');
});

});



