import { Component } from "@angular/core";

@Component({
  selector: "app-header",
  templateUrl: "./header.component.html",
  styleUrl: "./header.component.css",
})
export class HeaderComponent {
  public loggedIn: boolean = false;

  public username: string = "";
  public password: string = "";

  constructor() {}
}
