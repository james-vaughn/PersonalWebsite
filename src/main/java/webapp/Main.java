package main.java.webapp;

import spark.ModelAndView;
import spark.template.mustache.MustacheTemplateEngine;
import java.util.HashMap;
import static spark.Spark.*;

public class Main {
    public static void main(String[] args) {
        port(8000);
        staticFiles.location("/public");

        get("/", (request, response) -> {
            return new MustacheTemplateEngine().render(new ModelAndView(new HashMap<String, Object>(), "public/index.html"));
        });
    }
}
