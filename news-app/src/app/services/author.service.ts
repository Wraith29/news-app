import { Injectable } from "@angular/core";
import { CacheService } from "./cache.service";
import { environment } from "../../environments/environment";
import { Observable } from "rxjs";

@Injectable({
  providedIn: "root",
})
export class AuthorService {
  private _baseUrl: string;
  private _storageKey: string = "news-feed:authors";

  constructor(private _cache: CacheService) {
    this._baseUrl = environment.apiBaseUrl;
  }

  public getAll(): Observable<string[]> {
    const url = this._baseUrl + "authors";

    return this._cache.get<string[]>(url, this._storageKey);
  }
}
