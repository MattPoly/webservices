package org.polytech;

import io.spring.guides.gs_producing_web_service.Book;
import io.spring.guides.gs_producing_web_service.Books;
import org.springframework.stereotype.Component;

import javax.annotation.PostConstruct;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

@Component
public class BookRepository {
    private static final Books books = new Books();

    @PostConstruct
    public void initData() {
        Book first = new Book();
        first.setId(1);
        first.setName("A Game of Thrones");
        first.setNumberOfPages(647);
        first.setPublisher("Bantam Books");

        Book second = new Book();
        second.setId(2);
        second.setName("A Clash of Kings");
        first.setPublisher("Bantam Books");
        second.setNumberOfPages(768);

        Book third = new Book();
        third.setId(3);
        third.setName("A Storm of Swords");
        first.setPublisher("Bantam Books");
        third.setNumberOfPages(992);

        books.getElements().add(first);
        books.getElements().add(second);
        books.getElements().add(third);
    }

    public Book findBook(int id) {
        return books.getElements().get(id);
    }

    public Books findAll() {
        return books;
    }
}
