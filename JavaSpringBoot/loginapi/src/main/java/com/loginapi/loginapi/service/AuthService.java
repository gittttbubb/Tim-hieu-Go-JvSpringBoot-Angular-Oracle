package com.loginapi.loginapi.service;

import com.loginapi.loginapi.entity.User;
import com.loginapi.loginapi.repository.UserRepository;
import com.loginapi.loginapi.security.JwtUtil;
import lombok.RequiredArgsConstructor;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class AuthService {

    private final UserRepository repo;
    private final JwtUtil jwt;
    private final BCryptPasswordEncoder encoder = new BCryptPasswordEncoder();

    public String login(String username, String password) {
        User user = repo.findByUsername(username).orElseThrow(() -> new RuntimeException("account not found"));
        if (!encoder.matches(password,user.getPassword())) {
            throw new RuntimeException("wrong password");
        }
        return jwt.generateToken(user.getId());
    }

    public User getUserById(Integer id) {
        return repo.findById(id).orElseThrow();
    }

}