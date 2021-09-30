package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    status := []int{

        // http.StatusContinue,           // 100 // RFC 7231, 6.2.1 Get "http://localhost:8000": net/http: HTTP/1.x transport connection broken: unexpected EOF
        // http.StatusProcessing,         // 102 // RFC 2518, 10.1 Get "http://localhost:8000": net/http: HTTP/1.x transport connection broken: unexpected EOF
        // http.StatusMovedPermanently  , // 301 // RFC 7231, 6.4.2 Get "http://localhost:8000": 301 response missing Location header
        // http.StatusFound             , // 302 // RFC 7231, 6.4.3 Get "http://localhost:8000": 302 response missing Location header
        // http.StatusEarlyHints,         // 103 // RFC 8297 Get "http://localhost:8000": net/http: HTTP/1.x transport connection broken: unexpected EOF
        // http.StatusSeeOther,          // 303 // RFC 7231, 6.4.4 Get "http://localhost:8000": 303 response missing Location header

        http.StatusSwitchingProtocols, // 101 // RFC 7231, 6.2.2

        http.StatusOK,                   // 200 // RFC 7231, 6.3.1
        http.StatusCreated,              // 201 // RFC 7231, 6.3.2
        http.StatusAccepted,             // 202 // RFC 7231, 6.3.3
        http.StatusNonAuthoritativeInfo, // 203 // RFC 7231, 6.3.4
        http.StatusNoContent,            // 204 // RFC 7231, 6.3.5
        http.StatusResetContent,         // 205 // RFC 7231, 6.3.6
        http.StatusPartialContent,       // 206 // RFC 7233, 4.1
        http.StatusMultiStatus,          // 207 // RFC 4918, 11.1
        http.StatusAlreadyReported,      // 208 // RFC 5842, 7.1
        http.StatusIMUsed,               // 226 // RFC 3229, 10.4.1

        http.StatusMultipleChoices,   // 300 // RFC 7231, 6.4.1

        http.StatusNotModified,       // 304 // RFC 7232, 4.1
        http.StatusUseProxy,          // 305 // RFC 7231, 6.4.5
        http.StatusTemporaryRedirect, // 307 // RFC 7231, 6.4.7
        http.StatusPermanentRedirect, // 308 // RFC 7538, 3

        http.StatusBadRequest,                   // 400 // RFC 7231, 6.5.1
        http.StatusUnauthorized,                 // 401 // RFC 7235, 3.1
        http.StatusPaymentRequired,              // 402 // RFC 7231, 6.5.2
        http.StatusForbidden,                    // 403 // RFC 7231, 6.5.3
        http.StatusNotFound,                     // 404 // RFC 7231, 6.5.4
        http.StatusMethodNotAllowed,             // 405 // RFC 7231, 6.5.5
        http.StatusNotAcceptable,                // 406 // RFC 7231, 6.5.6
        http.StatusProxyAuthRequired,            // 407 // RFC 7235, 3.2
        http.StatusRequestTimeout,               // 408 // RFC 7231, 6.5.7
        http.StatusConflict,                     // 409 // RFC 7231, 6.5.8
        http.StatusGone,                         // 410 // RFC 7231, 6.5.9
        http.StatusLengthRequired,               // 411 // RFC 7231, 6.5.10
        http.StatusPreconditionFailed,           // 412 // RFC 7232, 4.2
        http.StatusRequestEntityTooLarge,        // 413 // RFC 7231, 6.5.11
        http.StatusRequestURITooLong,            // 414 // RFC 7231, 6.5.12
        http.StatusUnsupportedMediaType,         // 415 // RFC 7231, 6.5.13
        http.StatusRequestedRangeNotSatisfiable, // 416 // RFC 7233, 4.4
        http.StatusExpectationFailed,            // 417 // RFC 7231, 6.5.14
        http.StatusTeapot,                       // 418 // RFC 7168, 2.3.3
        http.StatusMisdirectedRequest,           // 421 // RFC 7540, 9.1.2
        http.StatusUnprocessableEntity,          // 422 // RFC 4918, 11.2
        http.StatusLocked,                       // 423 // RFC 4918, 11.3
        http.StatusFailedDependency,             // 424 // RFC 4918, 11.4
        http.StatusTooEarly,                     // 425 // RFC 8470, 5.2.
        http.StatusUpgradeRequired,              // 426 // RFC 7231, 6.5.15
        http.StatusPreconditionRequired,         // 428 // RFC 6585, 3
        http.StatusTooManyRequests,              // 429 // RFC 6585, 4
        http.StatusRequestHeaderFieldsTooLarge,  // 431 // RFC 6585, 5
        http.StatusUnavailableForLegalReasons,   // 451 // RFC 7725, 3

        http.StatusInternalServerError,           // 500 // RFC 7231, 6.6.1
        http.StatusNotImplemented,                // 501 // RFC 7231, 6.6.2
        http.StatusBadGateway,                    // 502 // RFC 7231, 6.6.3
        http.StatusServiceUnavailable,            // 503 // RFC 7231, 6.6.4
        http.StatusGatewayTimeout,                // 504 // RFC 7231, 6.6.5
        http.StatusHTTPVersionNotSupported,       // 505 // RFC 7231, 6.6.6
        http.StatusVariantAlsoNegotiates,         // 506 // RFC 2295, 8.1
        http.StatusInsufficientStorage,           // 507 // RFC 4918, 11.5
        http.StatusLoopDetected,                  // 508 // RFC 5842, 7.2
        http.StatusNotExtended,                   // 510 // RFC 2774, 7
        http.StatusNetworkAuthenticationRequired, // 511 // RFC 6585, 6
    }
    k := rand.Intn(len(status) - 1)
    w.WriteHeader(status[k])
    fmt.Println(w, status[k])
}

func main() {
    router := http.NewServeMux()
    server := &http.Server{
        Addr:         ":8000",
        Handler:      router,
        ReadTimeout:  5 * time.Second,
        //WriteTimeout: WriteTimeout + 10*time.Millisecond, //10ms Redundant time
        IdleTimeout:  15 * time.Second,
    }
    router.HandleFunc("/", indexHandler)
    server.ListenAndServe()
}
