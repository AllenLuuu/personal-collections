import { AlignmentType, Document, Packer, Paragraph } from "docx";
import { saveAs } from "file-saver";

const generateSource = (author: string, book: string) => {
  if (author && book) {
    return `——${author}《${book}》`;
  } else if (author) {
    return "——" + author;
  } else if (book) {
    return `——${book}》`;
  } else {
    return "";
  }
};

export const exportToDocx = (collections: CollectionType[]) => {
  const paragraphs: Paragraph[] = [];
  collections.forEach((collection) => {
    const { content, author, book } = collection;
    if (author || book) {
      paragraphs.push(
        new Paragraph({
          text: content,
          spacing: {
            after: 200,
          },
        }),
        new Paragraph({
          text: generateSource(author, book),
          alignment: AlignmentType.RIGHT,
          spacing: {
            after: 400,
          },
        })
      );
    } else {
      paragraphs.push(
        new Paragraph({
          text: content,
          spacing: {
            after: 400,
          },
        })
      );
    }
  });

  const doc = new Document({
    title: "摘抄导出",
    description: "导出自 Allenluuu 的摘抄",
    creator: "collections.allenluuu.com",
    sections: [
      {
        children: paragraphs,
      },
    ],
  });

  Packer.toBlob(doc).then((blob) => {
    console.log(blob);
    saveAs(blob, "摘抄导出.docx");
  });
};

export const exportToMarkdown = (collections: CollectionType[]) => {
  const markdown = collections
    .map((collection) => {
      const { content, author, book } = collection;
      if (author || book) {
        return `${content}\n<p align="right">${generateSource(
          author,
          book
        )}</p>\n`;
      } else {
        return `${content}\n`;
      }
    })
    .join("\n");
  const blob = new Blob([markdown], { type: "text/markdown" });
  saveAs(blob, "摘抄导出.md");
};
