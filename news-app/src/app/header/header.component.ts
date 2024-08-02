import { Component } from "@angular/core";
import { AuthService } from "../services/auth.service";

@Component({
  selector: "app-header",
  templateUrl: "./header.component.html",
  styleUrl: "./header.component.css",
})
export class HeaderComponent {
  public loginDialogVisible: boolean = false;
  public loggedIn: boolean = false;

  public username: string = "";
  public password: string = "";

  constructor(private _authService: AuthService) {
    this._updateDetails();
  }

  public openDialog(): void {
    this.loginDialogVisible = true;
  }

  public cancelDialog(): void {
    this.loginDialogVisible = false;
    this.username = "";
    this.password = "";
  }

  public login(): void {
    this._authService.login(this.username, this.password);
    this._updateDetails();
    this.loginDialogVisible = false;
  }

  private _updateDetails(): void {
    const username = this._authService.loggedInAs();

    this.loggedIn = username !== null;
    if (username !== null) this.username = username;
  }
}
