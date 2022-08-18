package com.fabrictemplate.template.dto;


import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor

@JsonSerialize
public class SocialMediaAccounts {
    private String websiteUrl;
    private String dateCreated;
}
