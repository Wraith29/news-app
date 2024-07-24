import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { map, Observable, Subscriber } from "rxjs";

interface CacheableRequest<T> {
  value: T;
  hash: string;
}

@Injectable({
  providedIn: "root",
})
export class CacheService {
  constructor(private _http: HttpClient) {}

  public get<T>(url: string, key: string): Observable<T> {
    const existingData = sessionStorage.getItem(key);

    if (existingData === null || existingData === "" || existingData === "{}") {
      return this._http.get<CacheableRequest<T>>(url).pipe(
        map((res: CacheableRequest<T>) => {
          this._writeToCache(key, res);

          return res.value;
        }),
      );
    }

    const data: CacheableRequest<T> = JSON.parse(existingData);

    return new Observable((sub: Subscriber<T>) => {
      sub.next(data.value);

      this._http
        .get<CacheableRequest<T>>(url)
        .pipe(
          map((res: CacheableRequest<T>) => {
            if (res.hash !== data.hash) {
              this._writeToCache(key, res);
              sub.next(res.value);
            }
          }),
        )
        .subscribe();
    });
  }

  private _writeToCache<T>(key: string, response: CacheableRequest<T>): void {
    sessionStorage.setItem(key, JSON.stringify(response));
  }
}
