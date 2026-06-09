package com.loginapi.loginapi.controller;

import com.loginapi.loginapi.dto.*;
import com.loginapi.loginapi.entity.User;
import com.loginapi.loginapi.service.AuthService;
import jakarta.servlet.http.HttpServletRequest;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;
import java.util.Map;

@RestController
@RequestMapping("/api")
//Cho phép bỏ qua CORS
@CrossOrigin(origins="http://localhost:4200")

@RequiredArgsConstructor
public class AuthController {
    private final AuthService service;
    @PostMapping("/auth/login")
    public LoginResponse login(@RequestBody LoginRequest req) {
        return new LoginResponse(service.login(
                        req.getUsername(),
                        req.getPassword()
                )
        );
    }

    @GetMapping("/users/me")
    public Map<String,Object> me(HttpServletRequest req) {
        Integer id = (Integer) req.getAttribute("userId");
        User user = service.getUserById(id);
        return Map.of("id", user.getId(),
                "username",user.getUsername(),
                "role", user.getRole()
        );
    }

}