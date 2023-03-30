import { Component } from '@angular/core';

@Component({
  selector: 'b-component',
  template: `
    <form class="box">

    <div class="field">
    <label class="label">First Name</label>
    <div class="control">
      <input class="input" type="first name" placeholder="ex. John">
    </div>
  </div>

  <div class="field">
    <label class="label">Last Name</label>
    <div class="control">
      <input class="input" type="last name" placeholder="ex. Doe">
    </div>
  </div>

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

  <button class="button is-primary">Create Account</button>


</form>

  `,
  styles: [`
     `]
})

export class BComponent {

}
