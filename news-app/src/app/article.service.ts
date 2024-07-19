import { Injectable } from '@angular/core';
import { Article } from './types/article';
import { Observable } from 'rxjs';
import { environment } from '../environments/environment';
import { CacheService } from './cache.service';

@Injectable({
  providedIn: 'root'
})
export class ArticleService {
  private _baseUrl: string;
  private _articleStorageKey: string = "articles";

  constructor(private _cache:CacheService) {
    this._baseUrl = environment.apiBaseUrl;
  }

  public getAllArticles(): Observable<Article[]> {
    const url = this._buildUrl("articles");

    return this._cache.get<Article[]>(url, this._articleStorageKey);

    // Nothing found in session storage, make request, write to session, pass values
    // const existingResponse = sessionStorage.getItem(this._sessionStorageKey);

    // if (existingResponse === null) {
    //   return this._http.get<ArticleResponse>(url).pipe(map((res: ArticleResponse) => {
    //     this._writeToCache(res);
    //     return res.articles ?? [];
    //   }));
    // }

    // const articleData: ArticleResponse = JSON.parse(existingResponse);

    // return new Observable((sub: Subscriber<Article[]>) => {
    //   sub.next(articleData.articles ?? []);

    //   this._http.get<ArticleResponse>(url).pipe(map((res: ArticleResponse) => {
    //     if (res.hash !== articleData.hash) {
    //       this._writeToCache(res);

    //       sub.next(res.articles ?? []);
    //     }
    //   })).subscribe();
    // });
  }

  private _buildUrl(endpoint: string): string {
    return this._baseUrl + endpoint;
  }

}
