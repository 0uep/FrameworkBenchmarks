package hello.web;

import java.io.*;
import java.util.*;

import javax.servlet.http.*;

import org.springframework.http.*;
import org.springframework.http.converter.*;
import org.springframework.http.converter.json.*;
import org.springframework.http.server.*;
import org.springframework.stereotype.*;
import org.springframework.web.bind.annotation.*;

/**
 * Handles requests for the application home page.
 */
@Controller
public class HelloJsonController
{

  @RequestMapping(value = "/json", produces = "application/json")
  @ResponseBody
  public Map<String, String> json()
  {
    Map<String, String> map = new HashMap<String, String>();
    map.put("message", "Hello, world");
    return map;
  }
}
