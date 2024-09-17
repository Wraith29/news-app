import {
  HttpEvent,
  HttpEventType,
  HttpHandlerFn,
  HttpRequest,
} from "@angular/common/http";
import { inject } from "@angular/core";
import { AuthService } from "../services/auth.service";
import { Observable, switchMap, tap } from "rxjs";

export function authInterceptor(
  req: HttpRequest<unknown>,
  next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
  const authService = inject(AuthService);

  // Don't add auth header to auth request
  if (req.url.endsWith("login") || req.url.endsWith("register")) {
    return next(req);
  }

  return authService.getAuthToken().pipe(
    switchMap((authToken: string) => {
      const updatedRequest = req.clone({
        setHeaders: {
          Authorization: authToken,
        },
      });

      return next(updatedRequest).pipe(
        tap((event: HttpEvent<unknown>) => {
          if (event.type === HttpEventType.Response) {
            console.log(req);

            const newToken = req.headers.get("Authorization");
            if (newToken !== null) {
              authService.updateToken(newToken);
            }
          }
        }),
      );
    }),
  );
}
