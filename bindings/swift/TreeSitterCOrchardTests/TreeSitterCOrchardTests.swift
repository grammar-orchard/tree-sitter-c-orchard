import XCTest
import SwiftTreeSitter
import TreeSitterCOrchard

final class TreeSitterCOrchardTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_c_orchard())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading C grammar")
    }
}
