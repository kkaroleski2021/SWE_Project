import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppComponent } from './app.component';

describe('AppComponent', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RouterTestingModule
      ],
      declarations: [
        AppComponent
      ],
    }).compileComponents();
  });
  

  it('should create the app', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app).toBeTruthy();
  });

  it(`should have as title 'hello-world'`, () => {
    const fixture = TestBed.createComponent(AppComponent);
    const app = fixture.componentInstance;
    expect(app.title).toEqual('hello-world');
  });

  it('should render title', () => {
    const fixture = TestBed.createComponent(AppComponent);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('.content span')?.textContent).toContain('hello-world app is running!');
  });

  it('should set image path as expected', () => {
    const fixture = TestBed.createComponent(AppComponent);
    const ele = fixture.debugElement.nativeElement.querySelectorAll('img');
    expect(ele[1]['src']).toContain('assets/images/stuffedAnimalChair.jpg'); // if you dont have `id` in `img` , you need to provide index as `ele[index]`
    expect(ele[2]['src']).toContain('assets/images/clamshellVase.jpg');
    expect(ele[3]['src']).toContain('assets/images/grandmasCouch.jpg');
    expect(ele[4]['src']).toContain('assets/images/clamshellVase.jpg');
});
});
