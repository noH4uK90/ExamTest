//
//  Question.swift
//  iOS
//
//  Created by Иван Спирин on 2/12/24.
//

import Foundation

struct Question: Codable, Identifiable {
    let id: Int
    let text: String
    let scoreId: Int
    let answers: [Answer]
}
