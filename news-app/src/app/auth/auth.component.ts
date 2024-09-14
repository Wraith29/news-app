import { Component, OnInit } from "@angular/core";
import { AuthService } from "../services/auth.service";
import { Router } from "@angular/router";
import { HttpErrorResponse } from "@angular/common/http";
import { AuthResponse } from "../types/authresponse";
import { AUTHTOKEN_KEY, USERNAME_KEY } from "../types/storage";

enum AuthTab {
  Login = "Login",
  Register = "Register",
}

@Component({
  selector: "app-auth",
  templateUrl: "./auth.component.html",
  styleUrl: "./auth.component.css",
})
export class AuthComponent implements OnInit {
  public tab: AuthTab = AuthTab.Login;
  public username: string = "";
  public password: string = "";
  public error: string = "";

  constructor(
    private _authService: AuthService,
    private _router: Router,
  ) {}

  ngOnInit(): void {
    if (this._authService.isLoggedIn()) {
      this._router.navigate(["profile"]);
    }
  }

  public isLogin(): boolean {
    return this.tab === AuthTab.Login;
  }

  public login(): void {
    this._authService.login(this.username, this.password).subscribe({
      next: (res: AuthResponse) => {
        localStorage.setItem(AUTHTOKEN_KEY, res.authToken);
        localStorage.setItem(USERNAME_KEY, this.username);

        this._router.navigate([""]);

        this._authService.loggedIn.emit(true);
      },
      error: (err: HttpErrorResponse) => {
        this.error = err.error.error;
      },
    });
  }
}
