import { Component, OnInit } from "@angular/core";
import { AuthService } from "../services/auth.service";
import { Router } from "@angular/router";

@Component({
  selector: "app-profile",
  templateUrl: "./profile.component.html",
  styleUrl: "./profile.component.css",
})
export class ProfileComponent implements OnInit {
  constructor(
    private _authService: AuthService,
    private _router: Router,
  ) {}

  ngOnInit(): void {
    if (!this._authService.isLoggedIn()) {
      this._router.navigate(["login"]);
    }
  }
}
