//
//  Answer.swift
//  iOS
//
//  Created by Иван Спирин on 2/12/24.
//

import Foundation

struct Answer: Codable, Identifiable {
    let id: Int
    let text: String
    let isRight: Bool
}
