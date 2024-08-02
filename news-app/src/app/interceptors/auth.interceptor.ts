import { HttpEvent, HttpHandlerFn, HttpRequest } from "@angular/common/http";
import { inject } from "@angular/core";
import { AuthService } from "../services/auth.service";
import { Observable, switchMap } from "rxjs";

export function authInterceptor(
  req: HttpRequest<unknown>,
  next: HttpHandlerFn,
): Observable<HttpEvent<unknown>> {
  const authService = inject(AuthService);

  // Don't add auth header to auth request
  if (req.url.endsWith("login")) {
    return next(req);
  }

  return authService.getAuthToken().pipe(
    switchMap((authToken: string) => {
      const updatedRequest = req.clone({
        setHeaders: {
          Authorization: authToken,
        },
      });

      return next(updatedRequest);
    }),
  );
}
