use askama_axum::Template;
use axum::{routing::get, Router};
use tower_http::services::ServeDir;

#[tokio::main]
async fn main() {
    // build our application with a single route
    let app = Router::new()
        .route("/", get(index))
        .nest_service("/assets", ServeDir::new("assets"));

    // run our app with hyper, listening globally on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

#[derive(Template)]
#[template(path = "index.html")]
struct IndexTemplate<'idx_templ> {
    title: &'idx_templ str,
}

async fn index() -> IndexTemplate<'static> {
    IndexTemplate { title: "ayo" }
}
