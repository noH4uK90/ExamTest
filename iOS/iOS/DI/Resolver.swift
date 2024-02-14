//
//  Resolver.swift
//  iOS
//
//  Created by Иван Спирин on 2/12/24.
//

import Foundation
import Swinject

@propertyWrapper
struct Inject<Component> {
    let wrappedValue: Component
    init() {
        self.wrappedValue = Resolver.shared.resolve(Component.self)
    }
}

class Resolver {
    static let shared = Resolver()
    private let container = buildContainer()

    func resolve<T>(_ type: T.Type) -> T {
        container.resolve(T.self)!
    }
}

func buildContainer() -> Container {
    let container = Container()

    container.register(DataTransferProtocol.self) { _ in
        return DataTransferService()
    }
    container.register(NetworkProtocol.self) { _ in
        return NetworkService()
    }

    return container
}
