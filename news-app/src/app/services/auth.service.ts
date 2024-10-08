import { HttpClient } from "@angular/common/http";
import { EventEmitter, Injectable } from "@angular/core";
import { environment } from "../../environments/environment";
import { Observable, Subscriber } from "rxjs";
import { AuthResponse } from "../types/authresponse";
import { AUTHTOKEN_KEY, USERNAME_KEY } from "../types/storage";
import { jwtDecode } from "jwt-decode";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  private _baseUrl: string;

  public loggedIn = new EventEmitter<boolean>();

  constructor(private _http: HttpClient) {
    this._baseUrl = environment.apiBaseUrl;

    this.loggedIn.emit(this.isLoggedIn());
  }

  public login(username: string, password: string): Observable<AuthResponse> {
    const url = this._baseUrl + "login";

    return this._http.post<AuthResponse>(url, {
      username: username,
      password: password,
    });
  }

  public register(
    username: string,
    password: string,
  ): Observable<AuthResponse> {
    const url = this._baseUrl + "register";

    return this._http.post<AuthResponse>(url, {
      username: username,
      password: password,
    });
  }

  public updateToken(newToken: string): void {
    localStorage.setItem(AUTHTOKEN_KEY, newToken);
  }

  public logout(): void {
    localStorage.removeItem(AUTHTOKEN_KEY);
    localStorage.removeItem(USERNAME_KEY);

    this.loggedIn.emit(false);
  }

  public isLoggedIn(): boolean {
    const authToken = localStorage.getItem(AUTHTOKEN_KEY);

    if (authToken === null) return false;

    const decodedToken = jwtDecode(authToken);

    if (!decodedToken.exp) return false;

    const now = Math.floor(Date.now() / 1000);

    if (now > decodedToken.exp!) return false;

    return true;
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
