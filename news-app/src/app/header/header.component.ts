import { Component, OnInit } from "@angular/core";
import { AuthService } from "../services/auth.service";

@Component({
  selector: "app-header",
  templateUrl: "./header.component.html",
  styleUrl: "./header.component.css",
})
export class HeaderComponent implements OnInit {
  public loggedIn: boolean = false;

  constructor(private _authService: AuthService) {}

  ngOnInit(): void {
    this._authService.loggedIn.subscribe({
      next: (isLoggedIn: boolean) => {
        this.loggedIn = isLoggedIn;
      },
    });
  }
}
