import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { environment } from "../../environments/environment";
import { Observable, Subscriber } from "rxjs";
import { AuthResponse } from "../types/authresponse";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  private _baseUrl: string;

  constructor(private _http: HttpClient) {
    this._baseUrl = environment.apiBaseUrl;
  }

  public login(username: string, password: string): void {
    const url = this._baseUrl + "login";

    this._http
      .post<AuthResponse>(url, {
        username: username,
        password: password,
      })
      .subscribe({
        next: (res: AuthResponse) => {
          localStorage.setItem("authToken", res.authToken);
          localStorage.setItem("loggedInAs", username);
        },
        error: (err: Error) => {
          console.error(err);
        },
      });
  }

  public logout(): void {
    localStorage.removeItem("username");
    localStorage.removeItem("authToken");
  }

  public loggedInAs(): string | null {
    return localStorage.getItem("loggedInAs");
  }

  /**
    @param username {string} The username
    @param password {string} The password (unhashed) 
  */
  public getAuthToken(): Observable<string> {
    return new Observable<string>((sub: Subscriber<string>) => {
      const authToken = localStorage.getItem("authToken");

      if (authToken !== null) {
        sub.next(authToken);
        sub.complete();
        return;
      }

      sub.error(new Error("No Auth Token found. Login to refresh"));
      sub.complete();
    });
  }
}
