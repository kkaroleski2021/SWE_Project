import { Component } from '@angular/core';

@Component({
  selector: 'b-component',
  template: `
    <form class="box">
  <div class="field">
    <label class="label">Email</label>
    <div class="control">
      <input class="input" type="email" placeholder="e.g. alex@example.com">
    </div>
  </div>
  <div class="field">
    <label class="label">Password</label>
    <div class="control">
      <input class="input" type="password" placeholder="********">
    </div>
  </div>
  
  <button>Log In</button>

  <p class="control">
  <button><a routerLink="Signup">
      Don't have an account? Sign up now!
      </a></button>
  </p>
</form>
  `,
  styles: [`
     `]
})

export class BComponent {

}

